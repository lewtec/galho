package database

import (
	"path/filepath"
	"testing"
)

func TestNewDatabaseModule(t *testing.T) {
	tests := []struct {
		name     string
		path     string
		wantName string
	}{
		{
			name:     "Standard Structure",
			path:     "internal/crm/db",
			wantName: "crm",
		},
		{
			name:     "Nested DB",
			path:     "internal/db/db",
			wantName: "db",
		},
		{
			name:     "App Fallback",
			path:     "internal/db/other",
			wantName: "app",
		},
		{
			name:     "Custom Path",
			path:     "internal/auth/custom_db",
			wantName: "auth",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Convert path to OS specific separator
			path := filepath.FromSlash(tt.path)
			m := NewDatabaseModule(path)
			if m.Name() != tt.wantName {
				t.Errorf("NewDatabaseModule(%q).Name() = %q, want %q", path, m.Name(), tt.wantName)
			}
		})
	}
}
