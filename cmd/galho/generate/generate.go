package generate

import (
	"github.com/lewtec/galho/pkg/core"
	"github.com/spf13/cobra"
)

var Command = &cobra.Command{
	Use:   "generate",
	Short: "Generate code or files",
}

// LoadCommands loads registered generate commands into the main generate command
func LoadCommands() {
	for _, cmd := range core.GetGenerateCommands() {
		Command.AddCommand(cmd)
	}
}
