package core

import (
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestGetProjectFindsMarker(t *testing.T) {
	root := t.TempDir()
	if err := os.WriteFile(filepath.Join(root, ".galho"), nil, 0o644); err != nil {
		t.Fatal(err)
	}
	// Nested cwd under the project root.
	nested := filepath.Join(root, "internal", "crm")
	if err := os.MkdirAll(nested, 0o755); err != nil {
		t.Fatal(err)
	}

	cwd, err := os.Getwd()
	if err != nil {
		t.Fatal(err)
	}
	if err := os.Chdir(nested); err != nil {
		t.Fatal(err)
	}
	t.Cleanup(func() { _ = os.Chdir(cwd) })

	p, err := GetProject()
	if err != nil {
		t.Fatalf("GetProject: %v", err)
	}
	got, err := filepath.EvalSymlinks(p.Dir())
	if err != nil {
		t.Fatal(err)
	}
	want, err := filepath.EvalSymlinks(root)
	if err != nil {
		t.Fatal(err)
	}
	if got != want {
		t.Fatalf("Dir() = %q, want %q", got, want)
	}
}

func TestGetProjectMissingMarker(t *testing.T) {
	root := t.TempDir()
	cwd, err := os.Getwd()
	if err != nil {
		t.Fatal(err)
	}
	if err := os.Chdir(root); err != nil {
		t.Fatal(err)
	}
	t.Cleanup(func() { _ = os.Chdir(cwd) })

	_, err = GetProject()
	if err == nil {
		t.Fatal("expected error when .galho is missing")
	}
	if !strings.Contains(err.Error(), "not a galho project") {
		t.Fatalf("error %q should mention not a galho project", err)
	}
}
