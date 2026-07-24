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
			return os.MkdirAll(destPath, 0o755)
		}

		info, err := d.Info()
		if err != nil {
			return err
		}
		perm := info.Mode().Perm()
		if perm == 0 {
			// Some FS implementations report no mode; default to rw-r--r--.
			perm = 0o644
		}

		srcFile, err := data.Open(path)
		if err != nil {
			return err
		}
		defer srcFile.Close()

		// OpenFile with perm so new files inherit the template mode (e.g. +x scripts).
		// Chmod after open so an existing destination is rewritten to the source mode.
		dstFile, err := os.OpenFile(destPath, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, perm)
		if err != nil {
			return err
		}

		if err := dstFile.Chmod(perm); err != nil {
			dstFile.Close()
			return err
		}

		if _, err := io.Copy(dstFile, srcFile); err != nil {
			dstFile.Close()
			return err
		}

		return dstFile.Close()
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
