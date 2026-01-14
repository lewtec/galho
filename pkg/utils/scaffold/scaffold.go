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
	return fs.WalkDir(data, ".", func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}

		// Calculate destination path, removing .tmpl extension if present
		destPath := filepath.Join(destination, path)
		if strings.HasSuffix(path, ".tmpl") {
			destPath = strings.TrimSuffix(destPath, ".tmpl")
		}

		// @SECURITY: Clean the path to prevent traversal attacks
		cleanDestPath := filepath.Clean(destPath)
		if !strings.HasPrefix(cleanDestPath, filepath.Clean(destination)+string(os.PathSeparator)) && cleanDestPath != filepath.Clean(destination) {
			return fmt.Errorf("path traversal attempt detected: %s", path)
		}
		destPath = cleanDestPath

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
