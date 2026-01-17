package database

import (
	"fmt"
	"path/filepath"

	"github.com/lewtec/galho/pkg/core"
)

type DatabaseModule struct {
	path string
	name string
}

func NewDatabaseModule(path string) *DatabaseModule {
	// Default name is the parent directory name
	// e.g. internal/crm/db -> crm
	dir := filepath.Dir(path)
	name := filepath.Base(dir)

	// Special case: if parent dir is "db" but current dir is NOT "db"
	// e.g. internal/db/something -> app
	// This maintains backward compatibility with previous logic
	if filepath.Base(path) != "db" && name == "db" {
		name = "app"
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

func (m *DatabaseModule) Name() string {
	return m.name
}

func (m *DatabaseModule) GenerateTasks() ([]core.Task, error) {
	taskName := fmt.Sprintf("gen:%s:db:sqlc", m.name)

	return []core.Task{
		{
			Name:        taskName,
			Description: fmt.Sprintf("Generate SQLC code for %s", m.name),
			Run:         "sqlc generate -f sqlc.yaml",
			Dir:         m.path,
		},
	}, nil
}
