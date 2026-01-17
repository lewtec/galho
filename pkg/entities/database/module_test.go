package database

import (
	"path/filepath"
	"testing"
)

func TestNewDatabaseModule(t *testing.T) {
	tests := []struct {
		path     string
		expected string
	}{
		{
			path:     "internal/crm/db",
			expected: "crm",
		},
		{
			path:     "internal/db",
			expected: "internal",
		},
		{
			path:     "internal/db/db",
			expected: "db",
		},
		{
			path:     "app/db",
			expected: "app",
		},
		{
			path:     "db/something",
			expected: "app",
		},
		{
			path:     "db",
			expected: ".", // Base("db") is "db", Base(Dir("db")) is ".", name=".". Base("db")=="db" -> name=Base(".")=".".
		},
	}

	for _, tt := range tests {
		t.Run(tt.path, func(t *testing.T) {
			// Normalize path for the current OS
			path := filepath.FromSlash(tt.path)
			mod := NewDatabaseModule(path)
			if mod.Name() != tt.expected {
				t.Errorf("NewDatabaseModule(%q).Name() = %q, want %q", path, mod.Name(), tt.expected)
			}
		})
	}
}
