package frontend

import (
	"path/filepath"
	"testing"
)

func TestNewFrontendModule(t *testing.T) {
	tests := []struct {
		name     string
		path     string
		wantName string
	}{
		{
			name:     "Standard Structure",
			path:     "internal/crm/frontend",
			wantName: "crm",
		},
		{
			name:     "Self Named",
			path:     "internal/frontend/frontend",
			wantName: "frontend",
		},
		{
			name:     "Custom Path",
			path:     "internal/auth/ui",
			wantName: "auth",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			path := filepath.FromSlash(tt.path)
			m := NewFrontendModule(path)
			if m.Name() != tt.wantName {
				t.Errorf("NewFrontendModule(%q).Name() = %q, want %q", path, m.Name(), tt.wantName)
			}
		})
	}
}
