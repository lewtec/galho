package scaffold

import (
	"io/fs"
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

func TestInstallFSPreservesExecutableMode(t *testing.T) {
	dest := t.TempDir()
	src := fstest.MapFS{
		"make_release": {
			Data: []byte("#!/usr/bin/env bash\necho ok\n"),
			Mode: 0o755,
		},
		"readme.txt": {
			Data: []byte("docs\n"),
			Mode: 0o644,
		},
	}
	if err := InstallFS(dest, src); err != nil {
		t.Fatal(err)
	}

	script := filepath.Join(dest, "make_release")
	info, err := os.Stat(script)
	if err != nil {
		t.Fatal(err)
	}
	if info.Mode().Perm()&0o111 == 0 {
		t.Fatalf("expected executable bits on make_release, got %o", info.Mode().Perm())
	}

	readme := filepath.Join(dest, "readme.txt")
	info, err = os.Stat(readme)
	if err != nil {
		t.Fatal(err)
	}
	if info.Mode().Perm()&0o111 != 0 {
		t.Fatalf("expected non-executable readme, got %o", info.Mode().Perm())
	}
}

func TestInstallFSRewritesExistingFileMode(t *testing.T) {
	dest := t.TempDir()
	path := filepath.Join(dest, "tool.sh")
	if err := os.WriteFile(path, []byte("old\n"), 0o644); err != nil {
		t.Fatal(err)
	}

	src := fstest.MapFS{
		"tool.sh": {
			Data: []byte("#!/bin/sh\n"),
			Mode: 0o755,
		},
	}
	if err := InstallFS(dest, src); err != nil {
		t.Fatal(err)
	}

	info, err := os.Stat(path)
	if err != nil {
		t.Fatal(err)
	}
	if got := info.Mode().Perm(); got&0o111 == 0 {
		t.Fatalf("expected +x after reinstall, got %o", got)
	}
	// Sanity: still a regular file mode we can inspect via fs.FileMode
	_ = fs.FileMode(0)
}
