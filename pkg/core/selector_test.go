package core

import (
	"testing"
)

type MockModule struct {
	path  string
	name  string
	mType string
}

func (m *MockModule) Type() string                   { return m.mType }
func (m *MockModule) Path() string                   { return m.path }
func (m *MockModule) Name() string                   { return m.name }
func (m *MockModule) GenerateTasks() ([]Task, error) { return nil, nil }

func TestStandardSelector_Select_Flag(t *testing.T) {
	selector := &StandardSelector{}
	modules := []Module{
		&MockModule{path: "internal/db", name: "db", mType: "database"},
		&MockModule{path: "internal/api", name: "api", mType: "database"},
	}

	// Test selecting by name (component)
	m, err := selector.Select(modules, "database", "api")
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	if m.Name() != "api" {
		t.Errorf("expected module 'api', got %s", m.Name())
	}

	// Test selecting by parent dir
	modules2 := []Module{
		&MockModule{path: "internal/crm/db", name: "db", mType: "database"},
	}
	m, err = selector.Select(modules2, "database", "crm")
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	if m.Path() != "internal/crm/db" {
		t.Errorf("expected module path 'internal/crm/db', got %s", m.Path())
	}
}

func TestStandardSelector_Select_Auto(t *testing.T) {
	selector := &StandardSelector{}
	modules := []Module{
		&MockModule{path: "internal/db", name: "db", mType: "database"},
	}

	// Test auto-selection
	m, err := selector.Select(modules, "database", "")
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	if m.Name() != "db" {
		t.Errorf("expected module 'db', got %s", m.Name())
	}
}
