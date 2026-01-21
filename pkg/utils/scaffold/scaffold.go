package scaffold

import (
	"io"
	"io/fs"
	"os"
	"path/filepath"
	"strings"
)

func InstallFS(destination string, data fs.FS) error {
	return fs.WalkDir(data, ".", func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}

		// Calculate destination path, removing .tmpl extension if present
		destPath := filepath.Join(destination, path)
		if strings.HasSuffix(path, ".tmpl") {
			destPath = strings.TrimSuffix(destPath, ".tmpl")
		}

		// Security: Prevent Zip Slip by ensuring destPath is inside destination
		// We use Abs to resolve relative paths and ensure prefix matching works
		absDest, err := filepath.Abs(destination)
		if err != nil {
			return err
		}
		absDestPath, err := filepath.Abs(destPath)
		if err != nil {
			return err
		}

		// Ensure the path is within the destination directory
		// We add a separator to ensure we don't match partial directory names (e.g. /tmp/foo vs /tmp/foobar)
		// But allow exact match if it is the root itself (though WalkDir usually starts with ".")
		if !strings.HasPrefix(absDestPath, absDest+string(os.PathSeparator)) && absDestPath != absDest {
			// Special case: if we are on the root of destination (e.g. "."), it matches.
			// But if we are somehow outside, return error.
			// Actually, if absDestPath == absDest, it means we are writing to the root.
			// WalkDir starts with "." so path is ".". destPath is destination.
			// So absDestPath == absDest is valid.
			// Any other file should have prefix absDest + separator.
			return os.ErrPermission // Or a more specific error
		}

		// If it's a directory, create it
		if d.IsDir() {
			return os.MkdirAll(destPath, 0755)
		}

		// Open source file
		srcFile, err := data.Open(path)
		if err != nil {
			return err
		}
		defer srcFile.Close()

		// Create destination file
		dstFile, err := os.Create(destPath)
		if err != nil {
			return err
		}
		defer dstFile.Close()

		// Copy contents
		if _, err := io.Copy(dstFile, srcFile); err != nil {
			return err
		}

		return nil
	})
}
