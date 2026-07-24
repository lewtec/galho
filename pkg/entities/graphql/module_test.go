package graphql

import "testing"

func TestNewGraphQLModuleName(t *testing.T) {
	m := NewGraphQLModule("internal/crm/api")
	if m.Name() != "crm" {
		t.Fatalf("Name() = %q, want crm", m.Name())
	}
	if m.Type() != "graphql" {
		t.Fatalf("Type() = %q", m.Type())
	}
}
