package graphql

import (
	"path/filepath"
	"testing"
)

func TestNewGraphQLModule(t *testing.T) {
	tests := []struct {
		name     string
		path     string
		wantName string
	}{
		{
			name:     "Standard Structure",
			path:     "internal/crm/api",
			wantName: "crm",
		},
		{
			name:     "Self Named",
			path:     "internal/api/api",
			wantName: "api",
		},
		{
			name:     "Custom Path",
			path:     "internal/auth/graphql",
			wantName: "auth",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			path := filepath.FromSlash(tt.path)
			m := NewGraphQLModule(path)
			if m.Name() != tt.wantName {
				t.Errorf("NewGraphQLModule(%q).Name() = %q, want %q", path, m.Name(), tt.wantName)
			}
		})
	}
}
