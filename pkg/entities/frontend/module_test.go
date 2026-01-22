package frontend

import (
	"path/filepath"
	"testing"
)

func TestNewFrontendModule(t *testing.T) {
	tests := []struct {
		path         string
		expectedName string
	}{
		{
			path:         filepath.Join("internal", "crm", "frontend"),
			expectedName: "crm",
		},
		{
			path:         filepath.Join("internal", "frontend", "frontend"),
			expectedName: "frontend",
		},
		{
			path:         filepath.Join("internal", "crm"),
			expectedName: "internal",
		},
	}

	for _, tt := range tests {
		t.Run(tt.path, func(t *testing.T) {
			m := NewFrontendModule(tt.path)
			if m.Name() != tt.expectedName {
				t.Errorf("NewFrontendModule(%q).Name() = %q, want %q", tt.path, m.Name(), tt.expectedName)
			}
		})
	}
}
