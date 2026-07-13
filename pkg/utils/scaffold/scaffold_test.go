package scaffold

import (
	"os"
	"path/filepath"
	"testing"
	"testing/fstest"
)

func TestSafeJoinRejectsEscape(t *testing.T) {
	dest := t.TempDir()
	if _, err := safeJoin(dest, "../outside"); err == nil {
		t.Fatal("expected escape to be rejected")
	}
	if _, err := safeJoin(dest, "ok/file.txt"); err != nil {
		t.Fatalf("expected safe path, got %v", err)
	}
}

func TestInstallFSWritesFiles(t *testing.T) {
	dest := t.TempDir()
	simple := fstest.MapFS{
		"hello.txt": {Data: []byte("hi")},
	}
	if err := InstallFS(dest, simple); err != nil {
		t.Fatal(err)
	}
	b, err := os.ReadFile(filepath.Join(dest, "hello.txt"))
	if err != nil {
		t.Fatal(err)
	}
	if string(b) != "hi" {
		t.Fatalf("got %q", b)
	}
}

func TestInstallFSRejectsEscapingPath(t *testing.T) {
	dest := t.TempDir()
	evil := fstest.MapFS{
		"../evil.txt": {Data: []byte("nope")},
	}
	if err := InstallFS(dest, evil); err == nil {
		t.Fatal("expected escape rejection")
	}
}
