package core

import "path/filepath"

// DeriveModuleName returns the friendly module name for a module directory path.
//
// Entity packages share the same convention: a module lives at
// .../<domain>/<leaf> (leaf is "db", "api", or "frontend") and the friendly
// name is the domain segment. This helper is the single source of that rule so
// database, graphql, and frontend constructors cannot drift.
//
// The "db" leaf keeps a historical edge case from NewDatabaseModule: when the
// parent directory is also named "db" and the module path itself is not "db"
// (e.g. "db/v1"), the name falls back to "app".
func DeriveModuleName(modulePath, leaf string) string {
	base := filepath.Base(modulePath)
	parent := filepath.Base(filepath.Dir(modulePath))

	if leaf == "db" && parent == "db" && base != "db" {
		return "app"
	}
	return parent
}
