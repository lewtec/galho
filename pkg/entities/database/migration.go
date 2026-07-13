package database

import (
	"fmt"
	"maps"
	"os"
	"path/filepath"
	"slices"
	"strings"
	"time"

	"github.com/lewtec/galho/pkg/core"

	"github.com/spf13/cobra"
)

func newMigrationCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "migration",
		Short: "Manage database migrations",
		Long:  "Create and list database migration files",
	}

	cmd.AddCommand(newMigrationCreateCommand())
	cmd.AddCommand(newMigrationListCommand())

	return cmd
}

func newMigrationCreateCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "create <name>",
		Short: "Create a new migration file",
		Long:  "Create a new timestamped migration file in the migrations directory",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			ctx, err := core.GetCommandContext(cmd)
			if err != nil {
				return err
			}

			migrationName := args[0]

			dbModule, ok := ctx.Module.(*DatabaseModule)
			if !ok {
				return fmt.Errorf("expected database module, got %s", ctx.Module.Type())
			}

			return createMigration(dbModule, migrationName)
		},
	}
}

func createMigration(module *DatabaseModule, name string) error {
	timestamp := time.Now().Format("20060102150405")
	filenameUp := fmt.Sprintf("%s_%s.up.sql", timestamp, name)
	filenameDown := fmt.Sprintf("%s_%s.down.sql", timestamp, name)

	migrationsDir := filepath.Join(module.Path(), "migrations")

	if err := os.MkdirAll(migrationsDir, 0755); err != nil {
		return fmt.Errorf("failed to create migrations directory: %w", err)
	}

	filepathUp := filepath.Join(migrationsDir, filenameUp)
	filepathDown := filepath.Join(migrationsDir, filenameDown)

	templateUp := fmt.Sprintf(`-- Migration: %s (up)
-- Created: %s

-- Write your "up" migration here
-- This migration will be applied when migrating forward

-- Example:
-- CREATE TABLE example (
--     id SERIAL PRIMARY KEY,
--     name TEXT NOT NULL,
--     created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
-- );
`, name, time.Now().Format("2006-01-02 15:04:05"))

	templateDown := fmt.Sprintf(`-- Migration: %s (down)
-- Created: %s

-- Write your "down" migration here
-- This migration will be applied when rolling back

-- Example:
-- DROP TABLE IF EXISTS example;
`, name, time.Now().Format("2006-01-02 15:04:05"))

	if err := os.WriteFile(filepathUp, []byte(templateUp), 0644); err != nil {
		return fmt.Errorf("failed to write up migration file: %w", err)
	}

	if err := os.WriteFile(filepathDown, []byte(templateDown), 0644); err != nil {
		return fmt.Errorf("failed to write down migration file: %w", err)
	}

	fmt.Printf("Created migrations:\n")
	fmt.Printf("  %s\n", filepathUp)
	fmt.Printf("  %s\n", filepathDown)

	return nil
}

func newMigrationListCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "list",
		Short: "List all migrations",
		Long:  "List all migration files in the migrations directory",
		RunE: func(cmd *cobra.Command, args []string) error {
			ctx, err := core.GetCommandContext(cmd)
			if err != nil {
				return err
			}

			dbModule, ok := ctx.Module.(*DatabaseModule)
			if !ok {
				return fmt.Errorf("expected database module, got %s", ctx.Module.Type())
			}

			return listMigrations(dbModule)
		},
	}
}

type migrationEntry struct {
	hasUp   bool
	hasDown bool
	modTime time.Time
}

func listMigrations(module *DatabaseModule) error {
	migrationsDir := filepath.Join(module.Path(), "migrations")

	entries, err := os.ReadDir(migrationsDir)
	if err != nil {
		if os.IsNotExist(err) {
			fmt.Printf("No migrations directory found in %s\n", module.Path())
			return nil
		}
		return fmt.Errorf("failed to read migrations: %w", err)
	}

	migrations := make(map[string]migrationEntry)

	for _, entry := range entries {
		if entry.IsDir() || filepath.Ext(entry.Name()) != ".sql" {
			continue
		}

		name := entry.Name()
		info, err := entry.Info()
		if err != nil {
			return fmt.Errorf("failed to stat migration %s: %w", name, err)
		}

		var baseName string
		var isUp, isDown bool
		switch {
		case strings.HasSuffix(name, ".up.sql"):
			baseName = strings.TrimSuffix(name, ".up.sql")
			isUp = true
		case strings.HasSuffix(name, ".down.sql"):
			baseName = strings.TrimSuffix(name, ".down.sql")
			isDown = true
		default:
			// Legacy single-file format without .up/.down
			baseName = strings.TrimSuffix(name, ".sql")
			isUp = true
		}

		m := migrations[baseName]
		if isUp {
			m.hasUp = true
		}
		if isDown {
			m.hasDown = true
		}
		if m.modTime.IsZero() || info.ModTime().After(m.modTime) {
			m.modTime = info.ModTime()
		}
		migrations[baseName] = m
	}

	fmt.Printf("Migrations in %s:\n\n", module.Name())

	if len(migrations) == 0 {
		fmt.Println("  No migrations found")
		return nil
	}

	keys := slices.Sorted(maps.Keys(migrations))

	for _, baseName := range keys {
		m := migrations[baseName]
		status := ""
		switch {
		case m.hasUp && m.hasDown:
			status = "[up/down]"
		case m.hasUp:
			status = "[up only]"
		case m.hasDown:
			status = "[down only]"
		}

		fmt.Printf("  %s %s  (%s)\n", baseName, status, m.modTime.Format("2006-01-02 15:04"))
	}

	fmt.Printf("\nTotal: %d migration(s)\n", len(migrations))

	return nil
}
