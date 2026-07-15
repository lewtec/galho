package core

import "testing"

// stubModule is a minimal Module for name-matching tests.
type stubModule struct {
	typ  string
	path string
	name string
}

func (m stubModule) Type() string                  { return m.typ }
func (m stubModule) Path() string                  { return m.path }
func (m stubModule) Name() string                  { return m.name }
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
		{"crm", true},       // Module.Name()
		{"db", true},        // basename
		{"internal", true},  // path component
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
	got := findModuleByName(mods, "crm")
	if got == nil || got.Name() != "crm" {
		t.Fatalf("findModuleByName(crm) = %v, want crm module", got)
	}
	if findModuleByName(mods, "missing") != nil {
		t.Fatal("expected nil for missing name")
	}
}
