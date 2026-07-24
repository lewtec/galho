package database

import "testing"

func TestNewDatabaseModuleName(t *testing.T) {
	m := NewDatabaseModule("internal/crm/db")
	if m.Name() != "crm" {
		t.Fatalf("Name() = %q, want crm", m.Name())
	}
	if m.Path() != "internal/crm/db" {
		t.Fatalf("Path() = %q", m.Path())
	}
	if m.Type() != "database" {
		t.Fatalf("Type() = %q", m.Type())
	}
}
