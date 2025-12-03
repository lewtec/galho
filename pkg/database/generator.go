package database

import (
	"fmt"
	"os"
	"path/filepath"
)

const sqlcTemplate = `version: "2"
sql:
  - schema: "migrations"
    queries: "queries.sql"
    engine: "postgresql"
    gen:
      go:
        package: "db"
        out: "."
        sql_package: "pgx/v5"
`

const queriesTemplate = `-- name: GetUser :one
SELECT * FROM users
WHERE id = $1 LIMIT 1;
`

const initialMigrationTemplate = `CREATE TABLE users (
  id   BIGSERIAL PRIMARY KEY,
  name text      NOT NULL
);
`

func Generate(path string) error {
	// Ensure directory exists
	if err := os.MkdirAll(path, 0755); err != nil {
		return err
	}

	// Create sqlc.yaml
	if err := os.WriteFile(filepath.Join(path, "sqlc.yaml"), []byte(sqlcTemplate), 0644); err != nil {
		return err
	}

	// Create queries.sql
	if err := os.WriteFile(filepath.Join(path, "queries.sql"), []byte(queriesTemplate), 0644); err != nil {
		return err
	}

	// Create migrations directory
	migrationsDir := filepath.Join(path, "migrations")
	if err := os.MkdirAll(migrationsDir, 0755); err != nil {
		return err
	}

	// Create initial migration
	if err := os.WriteFile(filepath.Join(migrationsDir, "001_initial.sql"), []byte(initialMigrationTemplate), 0644); err != nil {
		return err
	}

	fmt.Printf("Database module generated at %s\n", path)
	return nil
}
