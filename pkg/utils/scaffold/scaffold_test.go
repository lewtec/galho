package scaffold

import (
	"io/fs"
	"os"
	"path/filepath"
	"testing"
	"testing/fstest"
)

func TestInstallFS(t *testing.T) {
	// Create a temporary directory for testing
	tmpDir, err := os.MkdirTemp("", "test-install-fs")
	if err != nil {
		t.Fatalf("Failed to create temp dir: %v", err)
	}
	defer os.RemoveAll(tmpDir)

	// Create a test file system
	testFS := fstest.MapFS{
		"test.txt": {Data: []byte("hello")},
	}

	// Test case 1: Valid installation
	err = InstallFS(tmpDir, testFS)
	if err != nil {
		t.Errorf("InstallFS failed with valid path: %v", err)
	}

	// Check if the file was created
	_, err = os.Stat(filepath.Join(tmpDir, "test.txt"))
	if os.IsNotExist(err) {
		t.Errorf("File was not created in the destination directory")
	}

	// Test case 2: Path traversal attempt
	maliciousFS := fstest.MapFS{
		"../test.txt": {Data: []byte("malicious")},
	}

	err = InstallFS(tmpDir, maliciousFS)
	if err == nil {
		t.Errorf("InstallFS did not return an error for path traversal attempt")
	}
	if err != fs.ErrInvalid {
		t.Errorf("InstallFS did not return the expected error for path traversal attempt, got: %v", err)
	}

	// Check if the file was not created outside the destination directory
	// Note: We check in the parent of tmpDir for the malicious file
	_, err = os.Stat(filepath.Join(tmpDir, "..", "test.txt"))
	if !os.IsNotExist(err) {
		// Cleanup the maliciously created file if it exists
		os.Remove(filepath.Join(tmpDir, "..", "test.txt"))
		t.Errorf("File was created outside the destination directory")
	}
}
