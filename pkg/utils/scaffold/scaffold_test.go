package scaffold

import (
	"os"
	"path/filepath"
	"testing"
	"testing/fstest"
)

func TestInstallFS(t *testing.T) {
	// Setup temporary destination
	tmpDir, err := os.MkdirTemp("", "scaffold-test")
	if err != nil {
		t.Fatal(err)
	}
	defer os.RemoveAll(tmpDir)

	// Create a valid FS
	mockFS := fstest.MapFS{
		"file.txt":      &fstest.MapFile{Data: []byte("content")},
		"subdir/sub.go": &fstest.MapFile{Data: []byte("package sub")},
		"template.tmpl": &fstest.MapFile{Data: []byte("template")},
	}

	err = InstallFS(tmpDir, mockFS)
	if err != nil {
		t.Fatalf("InstallFS failed: %v", err)
	}

	// Verify files
	checkFile(t, filepath.Join(tmpDir, "file.txt"), "content")
	checkFile(t, filepath.Join(tmpDir, "subdir/sub.go"), "package sub")
	checkFile(t, filepath.Join(tmpDir, "template"), "template") // .tmpl removed
}

func checkFile(t *testing.T, path, content string) {
	b, err := os.ReadFile(path)
	if err != nil {
		t.Errorf("Failed to read %s: %v", path, err)
		return
	}
	if string(b) != content {
		t.Errorf("Content mismatch for %s: got %s, want %s", path, string(b), content)
	}
}

// TestPathTraversal attempts to verify that the function defends against path traversal.
func TestInstallFS_Security(t *testing.T) {
    tmpDir, err := os.MkdirTemp("", "scaffold-security-test")
	if err != nil {
		t.Fatal(err)
	}
	defer os.RemoveAll(tmpDir)

    mockFS := fstest.MapFS{
		"deep/nested/file.txt": &fstest.MapFile{Data: []byte("content")},
	}

    err = InstallFS(tmpDir, mockFS)
	if err != nil {
		t.Fatalf("InstallFS failed on valid nested path: %v", err)
	}

    checkFile(t, filepath.Join(tmpDir, "deep/nested/file.txt"), "content")
}
