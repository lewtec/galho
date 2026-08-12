package database

import (
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestValidateMigrationName(t *testing.T) {
	cases := []struct {
		name    string
		wantErr bool
	}{
		{"add_users", false},
		{"Add-Users_2", false},
		{"", true},
		{"../evil", true},
		{"foo/bar", true},
		{"foo\\bar", true},
		{"has space", true},
		{"weird;rm", true},
		{"..", true},
		{"a..b", true},
	}
	for _, tc := range cases {
		err := validateMigrationName(tc.name)
		if tc.wantErr && err == nil {
			t.Errorf("validateMigrationName(%q): want error", tc.name)
		}
		if !tc.wantErr && err != nil {
			t.Errorf("validateMigrationName(%q): unexpected error: %v", tc.name, err)
		}
	}
}

func TestMigrationFilePathRejectsEscape(t *testing.T) {
	dir := t.TempDir()
	// timestamp_ + "foo/../../../outside" cleans outside migrationsDir
	if _, err := migrationFilePath(dir, "20260101000000_foo/../../../outside.up.sql"); err == nil {
		t.Fatal("expected escape rejection for .. segments after a slash")
	}
	// nested file under migrations is also rejected
	if _, err := migrationFilePath(dir, "20260101000000_foo/bar.up.sql"); err == nil {
		t.Fatal("expected nested path rejection")
	}
	ok, err := migrationFilePath(dir, "20260101000000_add_users.up.sql")
	if err != nil {
		t.Fatalf("safe name: %v", err)
	}
	if filepath.Dir(ok) != dir {
		t.Fatalf("expected file under %s, got %s", dir, ok)
	}
}

func TestCreateMigration_PathTraversal(t *testing.T) {
	root := t.TempDir()
	mod := NewDatabaseModule(filepath.Join(root, "db"))
	if err := os.MkdirAll(mod.Path(), 0755); err != nil {
		t.Fatal(err)
	}

	// Malicious names must not write outside migrations/
	for _, name := range []string{"../../evil", "foo/../../../outside", "a/b"} {
		if err := createMigration(mod, name); err == nil {
			t.Fatalf("createMigration(%q): expected error", name)
		}
	}

	// No stray files outside migrations
	migrationsDir := filepath.Join(mod.Path(), "migrations")
	if _, err := os.Stat(migrationsDir); err == nil {
		// directory may exist from MkdirAll before validation — ensure empty of evil
		entries, _ := os.ReadDir(migrationsDir)
		if len(entries) != 0 {
			t.Fatalf("migrations dir should be empty after rejected creates, got %v", entries)
		}
	}
	// parent of module must not gain evil.up.sql siblings from traversal
	if entries, err := os.ReadDir(root); err == nil {
		for _, e := range entries {
			if strings.Contains(e.Name(), "evil") || strings.Contains(e.Name(), "outside") {
				t.Fatalf("unexpected file outside module: %s", e.Name())
			}
		}
	}
}

func TestCreateMigration_WritesUpAndDown(t *testing.T) {
	root := t.TempDir()
	mod := NewDatabaseModule(filepath.Join(root, "db"))
	if err := os.MkdirAll(mod.Path(), 0755); err != nil {
		t.Fatal(err)
	}
	if err := createMigration(mod, "add_users"); err != nil {
		t.Fatal(err)
	}
	migrationsDir := filepath.Join(mod.Path(), "migrations")
	entries, err := os.ReadDir(migrationsDir)
	if err != nil {
		t.Fatal(err)
	}
	if len(entries) != 2 {
		t.Fatalf("want 2 files, got %d: %v", len(entries), entries)
	}
	var sawUp, sawDown bool
	for _, e := range entries {
		if strings.HasSuffix(e.Name(), "_add_users.up.sql") {
			sawUp = true
		}
		if strings.HasSuffix(e.Name(), "_add_users.down.sql") {
			sawDown = true
		}
		// must stay under migrationsDir
		if filepath.Dir(filepath.Join(migrationsDir, e.Name())) != migrationsDir {
			t.Fatalf("file escaped: %s", e.Name())
		}
	}
	if !sawUp || !sawDown {
		t.Fatalf("missing up/down files: %v", entries)
	}
}
