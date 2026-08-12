package core

import "testing"

func TestDeriveModuleName(t *testing.T) {
	cases := []struct {
		path, leaf, want string
	}{
		// Conventional monorepo layout
		{"internal/crm/db", "db", "crm"},
		{"internal/auth/db", "db", "auth"},
		{"internal/crm/api", "api", "crm"},
		{"internal/admin/frontend", "frontend", "admin"},

		// Absolute paths (finders pass these)
		{"/home/u/proj/internal/crm/db", "db", "crm"},
		{"/home/u/proj/internal/crm/api", "api", "crm"},

		// Database-only historical fallback: parent is "db", path not ending in "db"
		{"db/v1", "db", "app"},
		{"internal/db/v1", "db", "app"},

		// Path is itself the leaf under a domain-like parent
		{"internal/db", "db", "internal"},
		{"db", "db", "."},
	}
	for _, tc := range cases {
		got := DeriveModuleName(tc.path, tc.leaf)
		if got != tc.want {
			t.Errorf("DeriveModuleName(%q, %q) = %q, want %q", tc.path, tc.leaf, got, tc.want)
		}
	}
}
