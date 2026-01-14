package scaffold

import (
	"os"
	"path/filepath"
	"testing"
	"testing/fstest"
)

func TestInstallFS_PathTraversal(t *testing.T) {
	// Create a temporary directory for the destination
	destDir, err := os.MkdirTemp("", "test-installfs-")
	if err != nil {
		t.Fatalf("Failed to create temp dir: %v", err)
	}
	defer os.RemoveAll(destDir)

	// Create a malicious file system with a path traversal attempt
	maliciousFS := fstest.MapFS{
		"../../../../tmp/pwned.txt": {
			Data: []byte("malicious content"),
		},
	}

	// Attempt to install the malicious file system
	err = InstallFS(destDir, maliciousFS)
	if err == nil {
		t.Error("Expected an error for path traversal, but got nil")
	}

	// Verify that the malicious file was not created
	pwnedFilePath := filepath.Join("/tmp", "pwned.txt")
	if _, err := os.Stat(pwnedFilePath); !os.IsNotExist(err) {
		t.Errorf("Malicious file was created at %s", pwnedFilePath)
		// Clean up the malicious file if it was created
		os.Remove(pwnedFilePath)
	}
}

func TestInstallFS_Valid(t *testing.T) {
	// Create a temporary directory for the destination
	destDir, err := os.MkdirTemp("", "test-installfs-valid-")
	if err != nil {
		t.Fatalf("Failed to create temp dir: %v", err)
	}
	defer os.RemoveAll(destDir)

	// Create a valid file system
	validFS := fstest.MapFS{
		"dir/file.txt": {
			Data: []byte("valid content"),
		},
		"file.txt.tmpl": {
			Data: []byte("template content"),
		},
	}

	// Install the valid file system
	err = InstallFS(destDir, validFS)
	if err != nil {
		t.Fatalf("InstallFS failed for valid files: %v", err)
	}

	// Verify that the files were created correctly
	// Check dir/file.txt
	content, err := os.ReadFile(filepath.Join(destDir, "dir/file.txt"))
	if err != nil {
		t.Errorf("Failed to read created file 'dir/file.txt': %v", err)
	}
	if string(content) != "valid content" {
		t.Errorf("Expected content 'valid content', got '%s'", string(content))
	}

	// Check file.txt (from file.txt.tmpl)
	content, err = os.ReadFile(filepath.Join(destDir, "file.txt"))
	if err != nil {
		t.Errorf("Failed to read created file 'file.txt': %v", err)
	}
	if string(content) != "template content" {
		t.Errorf("Expected content 'template content', got '%s'", string(content))
	}
}
