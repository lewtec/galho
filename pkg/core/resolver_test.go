package core

import (
	"testing"
)

type mockModule struct {
	path string
}

func (m mockModule) Path() string { return m.path }
func (m mockModule) Type() string { return "" }
func (m mockModule) Name() string { return "" }
func (m mockModule) GenerateTasks() ([]Task, error) { return nil, nil }

func TestModuleMatchesName(t *testing.T) {
	tests := []struct {
		name       string
		modulePath string
		searchName string
		want       bool
	}{
		{"Direct match", "internal/crm/db", "db", true},
		{"Parent match", "internal/crm/db", "crm", true},
		{"Component match", "internal/crm/db", "internal", true},
		{"Partial match fail", "internal/crm/db", "cr", false},
		{"No match", "internal/crm/db", "auth", false},
		{"Root match", "db", "db", true},
		{"Multi-level parent", "a/b/c/d", "b", true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := mockModule{path: tt.modulePath}
			if got := moduleMatchesName(m, tt.searchName); got != tt.want {
				t.Errorf("moduleMatchesName(%q, %q) = %v, want %v", tt.modulePath, tt.searchName, got, tt.want)
			}
		})
	}
}
