package core

import (
	"fmt"

	"github.com/spf13/cobra"
)

var generateCommands = make(map[string]*cobra.Command)

// RegisterGenerateCommand registers a generate command for an entity type.
func RegisterGenerateCommand(name string, cmd *cobra.Command) {
	if _, ok := generateCommands[name]; ok {
		panic(fmt.Sprintf("generate command %s is being registered twice", name))
	}
	generateCommands[name] = cmd
}

// GetGenerateCommands returns all registered generate commands.
func GetGenerateCommands() map[string]*cobra.Command {
	return generateCommands
}
