package database

import (
	"github.com/lewtec/galho/pkg/core"

	"github.com/spf13/cobra"
)

func init() {
	// Create root "db" command
	dbCmd := &cobra.Command{
		Use:     "db",
		Aliases: []string{"database"},
		Short:   "Database operations",
		Long:    "Manage database modules, migrations, and introspection",
	}

	// Add subcommands
	dbCmd.AddCommand(newMigrationCommand())

	// Register with core
	core.RegisterEntityCommand(core.EntityCommand{
		Name:          "db",
		LongName:      "database",
		EntityType:    "database",
		Command:       dbCmd,
		RequireModule: true,
	})
}
