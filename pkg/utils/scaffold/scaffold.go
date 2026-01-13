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

		// Security: Ensure the resolved path is within the destination directory
		absDest, err := filepath.Abs(destination)
		if err != nil {
			return err
		}
		absPath, err := filepath.Abs(destPath)
		if err != nil {
			return err
		}
		if !strings.HasPrefix(absPath, absDest) {
			return fs.ErrInvalid
		}

		// If it's a directory, create it
		if d.IsDir() {
			return os.MkdirAll(destPath, 0755)
		}

		// It's a file, create parent directory if needed
		if err := os.MkdirAll(filepath.Dir(destPath), 0755); err != nil {
			return err
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
