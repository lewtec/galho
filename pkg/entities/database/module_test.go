package database

import (
	"path/filepath"
	"testing"
)

func TestNewDatabaseModule(t *testing.T) {
	tests := []struct {
		path         string
		expectedName string
	}{
		{
			path:         filepath.Join("internal", "crm", "db"),
			expectedName: "crm",
		},
		{
			path:         filepath.Join("internal", "auth", "db"),
			expectedName: "auth",
		},
		// Test the "weird" case where path doesn't end in db
		// Current behavior: uses parent directory name
		{
			path:         filepath.Join("internal", "crm"),
			expectedName: "internal",
		},
		// Test the "db/db" case
		{
			path:         filepath.Join("internal", "db", "db"),
			expectedName: "db",
		},
		// Test root db case
		{
			path:         filepath.Join("app", "db"),
			expectedName: "app",
		},
	}

	for _, tt := range tests {
		t.Run(tt.path, func(t *testing.T) {
			m := NewDatabaseModule(tt.path)
			if m.Name() != tt.expectedName {
				t.Errorf("NewDatabaseModule(%q).Name() = %q, want %q", tt.path, m.Name(), tt.expectedName)
			}
		})
	}
}
