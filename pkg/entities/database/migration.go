package database

import (
	"fmt"
	"os"
	"path/filepath"
	"time"

	"galho/pkg/core"

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
			// Get command context
			ctx, err := core.GetCommandContext(cmd)
			if err != nil {
				return err
			}

			migrationName := args[0]

			// Get database module
			dbModule, ok := ctx.Module.(*DatabaseModule)
			if !ok {
				return fmt.Errorf("expected database module, got %s", ctx.Module.Type())
			}

			// Create migration
			return createMigration(dbModule, migrationName)
		},
	}
}

func createMigration(module *DatabaseModule, name string) error {
	// Generate migration filename with timestamp
	timestamp := time.Now().Format("20060102150405")
	filenameUp := fmt.Sprintf("%s_%s.up.sql", timestamp, name)
	filenameDown := fmt.Sprintf("%s_%s.down.sql", timestamp, name)

	// Create migrations directory path
	migrationsDir := filepath.Join(module.Path(), "migrations")

	// Ensure migrations directory exists
	if err := os.MkdirAll(migrationsDir, 0755); err != nil {
		return fmt.Errorf("failed to create migrations directory: %w", err)
	}

	// Full paths to migration files
	filepathUp := filepath.Join(migrationsDir, filenameUp)
	filepathDown := filepath.Join(migrationsDir, filenameDown)

	// Migration template for "up" file
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

	// Migration template for "down" file
	templateDown := fmt.Sprintf(`-- Migration: %s (down)
-- Created: %s

-- Write your "down" migration here
-- This migration will be applied when rolling back

-- Example:
-- DROP TABLE IF EXISTS example;
`, name, time.Now().Format("2006-01-02 15:04:05"))

	// Write "up" migration file
	if err := os.WriteFile(filepathUp, []byte(templateUp), 0644); err != nil {
		return fmt.Errorf("failed to write up migration file: %w", err)
	}

	// Write "down" migration file
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
			// Get command context
			ctx, err := core.GetCommandContext(cmd)
			if err != nil {
				return err
			}

			// Get database module
			dbModule, ok := ctx.Module.(*DatabaseModule)
			if !ok {
				return fmt.Errorf("expected database module, got %s", ctx.Module.Type())
			}

			// List migrations
			return listMigrations(dbModule)
		},
	}
}

func listMigrations(module *DatabaseModule) error {
	migrationsDir := filepath.Join(module.Path(), "migrations")

	// Read migrations directory
	entries, err := os.ReadDir(migrationsDir)
	if err != nil {
		if os.IsNotExist(err) {
			fmt.Printf("No migrations directory found in %s\n", module.Path())
			return nil
		}
		return fmt.Errorf("failed to read migrations: %w", err)
	}

	// Filter and display SQL files
	fmt.Printf("Migrations in %s:\n\n", module.Name())

	count := 0
	for _, entry := range entries {
		if !entry.IsDir() && filepath.Ext(entry.Name()) == ".sql" {
			info, _ := entry.Info()
			fmt.Printf("  %s  (%s)\n", entry.Name(), info.ModTime().Format("2006-01-02 15:04"))
			count++
		}
	}

	if count == 0 {
		fmt.Println("  No migrations found")
	} else {
		fmt.Printf("\nTotal: %d migration file(s)\n", count)
	}

	return nil
}
