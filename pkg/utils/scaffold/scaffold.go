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
	absDest, err := filepath.Abs(destination)
	if err != nil {
		return fmt.Errorf("failed to resolve destination path: %w", err)
	}

	return fs.WalkDir(data, ".", func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}

		// Calculate destination path, removing .tmpl extension if present
		destPath := filepath.Join(destination, path)
		if strings.HasSuffix(path, ".tmpl") {
			destPath = strings.TrimSuffix(destPath, ".tmpl")
		}

		// Security: Prevent Zip Slip / Path Traversal
		// Ensure the final path is within the destination directory
		absFinal, err := filepath.Abs(destPath)
		if err != nil {
			return fmt.Errorf("failed to resolve path %s: %w", destPath, err)
		}

		if absFinal != absDest && !strings.HasPrefix(absFinal, absDest+string(os.PathSeparator)) {
			return fmt.Errorf("security: illegal file path %s attempts to escape destination %s", path, destination)
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
