package core

import (
	"strings"
	"testing"
)

// stubModule is a minimal Module for name-matching tests.
type stubModule struct {
	typ  string
	path string
	name string
}

func (m stubModule) Type() string                   { return m.typ }
func (m stubModule) Path() string                   { return m.path }
func (m stubModule) Name() string                   { return m.name }
func (m stubModule) GenerateTasks() ([]Task, error) { return nil, nil }

func TestModuleMatchesName(t *testing.T) {
	m := stubModule{
		typ:  "database",
		path: "internal/crm/db",
		name: "crm",
	}

	cases := []struct {
		query string
		want  bool
	}{
		{"crm", true},      // Module.Name()
		{"db", true},       // basename
		{"internal", true}, // path component
		{"other", false},
		{"", false},
	}
	for _, tc := range cases {
		got := moduleMatchesName(m, tc.query)
		if got != tc.want {
			t.Errorf("moduleMatchesName(%q) = %v, want %v", tc.query, got, tc.want)
		}
	}
}

func TestModuleMatchesNamePrefersFriendlyName(t *testing.T) {
	// Name is not a path component (e.g. fallback naming to "app").
	m := stubModule{
		typ:  "database",
		path: "internal/db",
		name: "app",
	}
	if !moduleMatchesName(m, "app") {
		t.Fatal("expected match on Module.Name() even when absent from path")
	}
	if moduleMatchesName(m, "appx") {
		t.Fatal("expected non-match for unrelated query")
	}
}

func TestFindModuleByName(t *testing.T) {
	mods := []Module{
		stubModule{path: "internal/auth/db", name: "auth"},
		stubModule{path: "internal/crm/db", name: "crm"},
	}
	got, err := findModuleByName(mods, "crm")
	if err != nil {
		t.Fatalf("findModuleByName(crm): %v", err)
	}
	if got.Name() != "crm" {
		t.Fatalf("findModuleByName(crm) = %v, want crm module", got)
	}
	if _, err := findModuleByName(mods, "missing"); err == nil {
		t.Fatal("expected error for missing name")
	} else if !strings.Contains(err.Error(), "not found") {
		t.Fatalf("error %q should say not found", err)
	}
}

func TestFindModuleByNameAmbiguousBasename(t *testing.T) {
	// Both modules share basename "db"; without disambiguation the first walker
	// result would win silently.
	mods := []Module{
		stubModule{path: "internal/auth/db", name: "auth"},
		stubModule{path: "internal/crm/db", name: "crm"},
	}
	_, err := findModuleByName(mods, "db")
	if err == nil {
		t.Fatal("expected ambiguous error for shared basename db")
	}
	if !strings.Contains(err.Error(), "ambiguous") {
		t.Fatalf("error %q should mention ambiguous", err)
	}
	if !strings.Contains(err.Error(), "internal/auth/db") || !strings.Contains(err.Error(), "internal/crm/db") {
		t.Fatalf("error %q should list matching paths", err)
	}
}

func TestFindModuleByNamePrefersUniqueFriendlyName(t *testing.T) {
	// Path-based match would hit both (shared "db"), but Name() uniquely selects.
	mods := []Module{
		stubModule{path: "internal/auth/db", name: "auth"},
		stubModule{path: "internal/crm/db", name: "crm"},
	}
	got, err := findModuleByName(mods, "auth")
	if err != nil {
		t.Fatalf("findModuleByName(auth): %v", err)
	}
	if got.Path() != "internal/auth/db" {
		t.Fatalf("got path %q, want internal/auth/db", got.Path())
	}
}

func TestFindModuleByNameUniquePathFallback(t *testing.T) {
	mods := []Module{
		stubModule{path: "internal/crm/db", name: "crm"},
	}
	got, err := findModuleByName(mods, "db")
	if err != nil {
		t.Fatalf("findModuleByName(db): %v", err)
	}
	if got.Name() != "crm" {
		t.Fatalf("got name %q, want crm", got.Name())
	}
}
