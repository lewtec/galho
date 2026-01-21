package core

import (
	"testing"
)

// MockModule implements Module interface for testing
type MockModule struct {
	path string
}

func (m *MockModule) Path() string { return m.path }
func (m *MockModule) Name() string { return "mock" }
func (m *MockModule) Type() string { return "mock" }
func (m *MockModule) GenerateTasks() ([]Task, error) { return nil, nil }

func TestModuleMatchesName(t *testing.T) {
	tests := []struct {
		name     string
		path     string
		match    string
		expected bool
	}{
		{"Exact match", "internal/crm/db", "db", true},
		{"Parent match", "internal/crm/db", "crm", true},
		{"Ancestor match", "internal/crm/db", "internal", true},
		{"No match", "internal/crm/db", "api", false},
		{"Partial match fail", "internal/crm/db", "cr", false},
		{"Empty match fail", "internal/crm/db", "", false},
		{"Windows path match", "internal/crm/db", "crm", true}, // logic uses ToSlash so it should work if we pass slash paths.
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &MockModule{path: tt.path}
			if got := moduleMatchesName(m, tt.match); got != tt.expected {
				t.Errorf("moduleMatchesName(%q, %q) = %v, want %v", tt.path, tt.match, got, tt.expected)
			}
		})
	}
}
