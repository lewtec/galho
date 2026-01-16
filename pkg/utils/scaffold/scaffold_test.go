
package scaffold

import (
	"io/fs"
	"os"
	"path/filepath"
	"testing"
	"testing/fstest"
)

// nonCompliantDirEntry is a mock fs.DirEntry for testing.
type nonCompliantDirEntry struct {
	name string
}

func (d *nonCompliantDirEntry) Name() string      { return d.name }
func (d *nonCompliantDirEntry) IsDir() bool       { return false }
func (d *nonCompliantDirEntry) Type() fs.FileMode { return 0 }
func (d *nonCompliantDirEntry) Info() (fs.FileInfo, error) { return nil, nil }

// nonCompliantFS is a mock fs.FS that produces a malicious path.
type nonCompliantFS struct {
	fs fstest.MapFS
}

func (mfs *nonCompliantFS) Open(name string) (fs.File, error) {
	return mfs.fs.Open(name)
}

// ReadDir returns a slice of DirEntry containing a malicious entry.
func (mfs *nonCompliantFS) ReadDir(name string) ([]fs.DirEntry, error) {
	if name == "." {
		return []fs.DirEntry{&nonCompliantDirEntry{name: "../evil.txt"}}, nil
	}
	return nil, fs.ErrNotExist
}

func TestInstallFS_Success(t *testing.T) {
	destDir, err := os.MkdirTemp("", "test-installfs-success")
	if err != nil {
		t.Fatalf("Failed to create temp dir: %v", err)
	}
	defer os.RemoveAll(destDir)

	testFS := fstest.MapFS{
		"dir/file.txt": {
			Data: []byte("hello world"),
		},
	}

	if err := InstallFS(destDir, testFS); err != nil {
		t.Fatalf("InstallFS failed: %v", err)
	}

	content, err := os.ReadFile(filepath.Join(destDir, "dir/file.txt"))
	if err != nil {
		t.Errorf("Failed to read created file: %v", err)
	}
	if string(content) != "hello world" {
		t.Errorf("Unexpected file content: got %q, want %q", string(content), "hello world")
	}
}

func TestInstallFS_PathTraversal(t *testing.T) {
	destDir, err := os.MkdirTemp("", "test-installfs-traversal")
	if err != nil {
		t.Fatalf("Failed to create temp dir: %v", err)
	}
	defer os.RemoveAll(destDir)

	maliciousFS := &nonCompliantFS{
		fs: fstest.MapFS{
			"../evil.txt": {
				Data: []byte("malicious content"),
			},
		},
	}

	err = InstallFS(destDir, maliciousFS)
	if err != fs.ErrPermission {
		t.Errorf("Expected fs.ErrPermission, but got: %v", err)
	}

	// Verify that the malicious file was not created outside the destination directory
	evilFilePath, _ := filepath.Abs(filepath.Join(destDir, "../evil.txt"))
	if _, err := os.Stat(evilFilePath); !os.IsNotExist(err) {
		t.Errorf("Malicious file was created at: %s", evilFilePath)
	}
}
