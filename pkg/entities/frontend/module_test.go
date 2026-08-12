package frontend

import "testing"

func TestNewFrontendModuleName(t *testing.T) {
	m := NewFrontendModule("internal/admin/frontend")
	if m.Name() != "admin" {
		t.Fatalf("Name() = %q, want admin", m.Name())
	}
	if m.Type() != "frontend" {
		t.Fatalf("Type() = %q", m.Type())
	}
}
