package core

import (
	"errors"
	"testing"
)

func TestNewScaffoldGenerateCommandInstallsPath(t *testing.T) {
	var got string
	cmd := NewScaffoldGenerateCommand("thing [path]", "Generate a thing", func(path string) error {
		got = path
		return nil
	})
	cmd.SetArgs([]string{"internal/crm/db"})
	if err := cmd.Execute(); err != nil {
		t.Fatalf("Execute: %v", err)
	}
	if got != "internal/crm/db" {
		t.Fatalf("install path = %q, want internal/crm/db", got)
	}
}

var errScaffoldGenerateBoom = errors.New("boom")

func TestNewScaffoldGenerateCommandPropagatesError(t *testing.T) {
	cmd := NewScaffoldGenerateCommand("thing [path]", "Generate a thing", func(path string) error {
		return errScaffoldGenerateBoom
	})
	cmd.SetArgs([]string{"x"})
	if err := cmd.Execute(); !errors.Is(err, errScaffoldGenerateBoom) {
		t.Fatalf("Execute error = %v, want %v", err, errScaffoldGenerateBoom)
	}
}

func TestNewScaffoldGenerateCommandRequiresPath(t *testing.T) {
	cmd := NewScaffoldGenerateCommand("thing [path]", "Generate a thing", func(path string) error {
		t.Fatal("install should not run without args")
		return nil
	})
	cmd.SetArgs(nil)
	if err := cmd.Execute(); err == nil {
		t.Fatal("expected error for missing path arg")
	}
}
