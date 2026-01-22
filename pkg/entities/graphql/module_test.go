package graphql

import (
	"path/filepath"
	"testing"
)

func TestNewGraphQLModule(t *testing.T) {
	tests := []struct {
		path         string
		expectedName string
	}{
		{
			path:         filepath.Join("internal", "crm", "api"),
			expectedName: "crm",
		},
		{
			path:         filepath.Join("internal", "api", "api"),
			expectedName: "api",
		},
		{
			path:         filepath.Join("internal", "crm"),
			expectedName: "internal",
		},
	}

	for _, tt := range tests {
		t.Run(tt.path, func(t *testing.T) {
			m := NewGraphQLModule(tt.path)
			if m.Name() != tt.expectedName {
				t.Errorf("NewGraphQLModule(%q).Name() = %q, want %q", tt.path, m.Name(), tt.expectedName)
			}
		})
	}
}
