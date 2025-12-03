package database

import (
	"fmt"
	"path/filepath"

	"galho/pkg/core"
)

type DatabaseModule struct {
	path string
	name string
}

func NewDatabaseModule(path string) *DatabaseModule {
	name := filepath.Base(filepath.Dir(path)) // Assuming structure internal/NAME/db/
	if name == "db" {
		// Fallback if path is just internal/db
		name = "app"
	}
	// Try to get the parent module name
	// e.g. internal/crm/db -> crm
	// internal/auth/db -> auth

	dir := filepath.Dir(path)
	if filepath.Base(path) == "db" {
		name = filepath.Base(dir)
	}

	return &DatabaseModule{
		path: path,
		name: name,
	}
}

func (m *DatabaseModule) Type() string {
	return "database"
}

func (m *DatabaseModule) Path() string {
	return m.path
}

func (m *DatabaseModule) GenerateTasks() ([]core.Task, error) {
	taskName := fmt.Sprintf("gen:%s:db:sqlc", m.name)
	sqlcConfigPath := filepath.Join(m.path, "sqlc.yaml")

	return []core.Task{
		{
			Name:        taskName,
			Description: fmt.Sprintf("Generate SQLC code for %s", m.name),
			Run:         fmt.Sprintf("sqlc generate -f %s", sqlcConfigPath),
			Dir:         m.path,
		},
	}, nil
}
