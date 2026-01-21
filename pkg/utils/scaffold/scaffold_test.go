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
		"file1.txt":     {Data: []byte("content1")},
		"dir/file2.txt": {Data: []byte("content2")},
		"template.tmpl": {Data: []byte("template content")},
	}

	// Test installation
	err = InstallFS(tmpDir, mockFS)
	if err != nil {
		t.Fatalf("InstallFS failed: %v", err)
	}

	// Verify files exist
	expectedFiles := []string{
		"file1.txt",
		"dir/file2.txt",
		"template", // .tmpl should be stripped
	}

	for _, f := range expectedFiles {
		path := filepath.Join(tmpDir, f)
		if _, err := os.Stat(path); os.IsNotExist(err) {
			t.Errorf("Expected file %s not found", path)
		}
	}

	// Check content of template (should be just copied, stripping extension only affects name)
	content, err := os.ReadFile(filepath.Join(tmpDir, "template"))
	if err != nil {
		t.Fatal(err)
	}
	if string(content) != "template content" {
		t.Errorf("Unexpected content: %s", string(content))
	}
}
