package core

import (
	"fmt"
	"maps"
	"slices"

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

// GetGenerateCommands returns registered generate commands ordered by registration name.
func GetGenerateCommands() []*cobra.Command {
	names := slices.Sorted(maps.Keys(generateCommands))
	out := make([]*cobra.Command, 0, len(names))
	for _, name := range names {
		out = append(out, generateCommands[name])
	}
	return out
}
