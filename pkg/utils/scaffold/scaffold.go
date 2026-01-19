package scaffold

import (
	"fmt"
	"io"
	"io/fs"
	"os"
	"path/filepath"
	"strings"
)

func InstallFS(destination string, data fs.FS) error {
	// SECURITY: Ensure destination is absolute to prevent relative path ambiguity
	absDest, err := filepath.Abs(destination)
	if err != nil {
		return err
	}

	return fs.WalkDir(data, ".", func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}

		// Calculate destination path, removing .tmpl extension if present
		destPath := filepath.Join(absDest, path)
		if strings.HasSuffix(path, ".tmpl") {
			destPath = strings.TrimSuffix(destPath, ".tmpl")
		}

		// SECURITY: Path Traversal Check (Zip Slip)
		if destPath != absDest && !strings.HasPrefix(destPath, absDest+string(os.PathSeparator)) {
			return fmt.Errorf("illegal file path: %s", destPath)
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
