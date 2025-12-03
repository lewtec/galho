package templates

import (
	"embed"
	"fmt"
	"io"
	"io/fs"
	"os"
	"path/filepath"
)

//go:embed all:database all:graphql all:frontend
var FS embed.FS

// CopyDir copies all files from the embedded FS directory src to the local file system directory dst.
func CopyDir(src string, dst string) error {
	return fs.WalkDir(FS, src, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}

		relPath, err := filepath.Rel(src, path)
		if err != nil {
			return err
		}

		targetPath := filepath.Join(dst, relPath)

		if d.IsDir() {
			if err := os.MkdirAll(targetPath, 0755); err != nil {
				return err
			}
			return nil
		}

		// Copy file content
		sourceFile, err := FS.Open(path)
		if err != nil {
			return err
		}
		defer sourceFile.Close()

		destFile, err := os.Create(targetPath)
		if err != nil {
			return err
		}
		defer destFile.Close()

		if _, err := io.Copy(destFile, sourceFile); err != nil {
			return err
		}

		fmt.Printf("Created %s\n", targetPath)
		return nil
	})
}
