package entities

import (
	"github.com/lewtec/galho/pkg/core"
	"github.com/spf13/cobra"
)

// AddEntityCommands adds registered entity commands to the root command.
func AddEntityCommands(rootCmd *cobra.Command) {
	for _, config := range core.GetEntityCommands() {
		rootCmd.AddCommand(config.Command)
	}
}
