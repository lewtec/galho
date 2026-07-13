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
	destination = filepath.Clean(destination)

	return fs.WalkDir(data, ".", func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}

		// Destination path; strip .tmpl so templates install without the suffix.
		rel := path
		if strings.HasSuffix(rel, ".tmpl") {
			rel = strings.TrimSuffix(rel, ".tmpl")
		}

		destPath, err := safeJoin(destination, rel)
		if err != nil {
			return err
		}

		// WalkDir visits directories before their contents.
		if d.IsDir() {
			return os.MkdirAll(destPath, 0755)
		}

		srcFile, err := data.Open(path)
		if err != nil {
			return err
		}
		defer srcFile.Close()

		dstFile, err := os.Create(destPath)
		if err != nil {
			return err
		}
		defer dstFile.Close()

		if _, err := io.Copy(dstFile, srcFile); err != nil {
			return err
		}

		return nil
	})
}

// safeJoin joins destination with a relative path and rejects escapes outside destination.
func safeJoin(destination, rel string) (string, error) {
	if rel == "." {
		return destination, nil
	}
	if filepath.IsAbs(rel) {
		return "", fmt.Errorf("scaffold: refusing absolute path %q", rel)
	}
	cleanRel := filepath.Clean(rel)
	if cleanRel == ".." || strings.HasPrefix(cleanRel, ".."+string(os.PathSeparator)) {
		return "", fmt.Errorf("scaffold: path escapes destination: %q", rel)
	}

	destPath := filepath.Join(destination, cleanRel)
	relToDest, err := filepath.Rel(destination, destPath)
	if err != nil {
		return "", fmt.Errorf("scaffold: invalid path %q: %w", rel, err)
	}
	if relToDest == ".." || strings.HasPrefix(relToDest, ".."+string(os.PathSeparator)) {
		return "", fmt.Errorf("scaffold: path escapes destination: %q", rel)
	}
	return destPath, nil
}
