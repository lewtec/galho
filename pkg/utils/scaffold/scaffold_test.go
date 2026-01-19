package scaffold

import (
	"os"
	"path/filepath"
	"testing"
	"testing/fstest"
)

func TestInstallFS(t *testing.T) {
	// Create a temporary directory for destination
	tmpDir, err := os.MkdirTemp("", "scaffold-test")
	if err != nil {
		t.Fatal(err)
	}
	defer os.RemoveAll(tmpDir)

	// Create a mock FS
	mockFS := fstest.MapFS{
		"file.txt": &fstest.MapFile{
			Data: []byte("hello world"),
		},
		"subdir/subfile.txt": &fstest.MapFile{
			Data: []byte("hello sub"),
		},
		"template.tmpl": &fstest.MapFile{
			Data: []byte("template content"),
		},
	}

	// Run InstallFS with absolute path (from MkdirTemp)
	err = InstallFS(tmpDir, mockFS)
	if err != nil {
		t.Fatalf("InstallFS failed with absolute path: %v", err)
	}

	// Verify files exist
	checkFile(t, filepath.Join(tmpDir, "file.txt"), "hello world")
	checkFile(t, filepath.Join(tmpDir, "subdir/subfile.txt"), "hello sub")
	checkFile(t, filepath.Join(tmpDir, "template"), "template content")

	// Test with relative path
	// We need to change directory to a temp location to safely test relative paths
	workDir, err := os.MkdirTemp("", "scaffold-work")
	if err != nil {
		t.Fatal(err)
	}
	defer os.RemoveAll(workDir)

	cwd, err := os.Getwd()
	if err != nil {
		t.Fatal(err)
	}

	if err := os.Chdir(workDir); err != nil {
		t.Fatal(err)
	}
	defer os.Chdir(cwd)

	relDest := "rel-dest"
	if err := os.Mkdir(relDest, 0755); err != nil {
		t.Fatal(err)
	}

	err = InstallFS(relDest, mockFS)
	if err != nil {
		t.Fatalf("InstallFS failed with relative path: %v", err)
	}

	// Check file in relative path
	absRelDest, _ := filepath.Abs(relDest)
	checkFile(t, filepath.Join(absRelDest, "file.txt"), "hello world")
}

func checkFile(t *testing.T, path, content string) {
	data, err := os.ReadFile(path)
	if err != nil {
		t.Errorf("Failed to read file %s: %v", path, err)
		return
	}
	if string(data) != content {
		t.Errorf("File %s content mismatch: got %q, want %q", path, string(data), content)
	}
}
