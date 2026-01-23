package database

import (
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/lewtec/galho/pkg/core"
)

func TestCreateMigration_PathTraversal(t *testing.T) {
	// Setup temp dir
	tmpDir, err := os.MkdirTemp("", "galho-test-*")
	if err != nil {
		t.Fatal(err)
	}
	defer os.RemoveAll(tmpDir)

	// Create a dummy module structure
	// /tmp/pkg/db
	modulePath := filepath.Join(tmpDir, "pkg", "db")
	err = os.MkdirAll(modulePath, 0755)
	if err != nil {
		t.Fatal(err)
	}

	module := NewDatabaseModule(modulePath)

	// Construct the command
	cmd := newMigrationCreateCommand()
	// Mock the context
	core.SetCommandContext(cmd, &core.CommandContext{
		Module: module,
	})

	// Attack payload: attempt to write to tmpDir root
	evilName := "../../../../../evil"

	// Run command
	// We expect this to FAIL with our validation error
	err = cmd.RunE(cmd, []string{evilName})
	if err == nil {
		t.Error("Vulnerability: Command succeeded with path traversal input")
	} else {
		if !strings.Contains(err.Error(), "invalid migration name") {
			t.Errorf("Unexpected error message: %v", err)
		}
	}

	// Double check that file does NOT exist
	files, err := os.ReadDir(tmpDir)
	if err != nil {
		t.Fatal(err)
	}

	found := false
	for _, f := range files {
		if !f.IsDir() && strings.HasSuffix(f.Name(), "evil.up.sql") {
			found = true
			break
		}
	}

	if found {
		t.Errorf("Vulnerability CONFIRMED: File created in tmpDir despite error: %s", tmpDir)
	}
}
