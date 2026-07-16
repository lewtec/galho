package core

import (
	"fmt"
	"maps"
	"slices"

	"github.com/spf13/cobra"
)

// EntityCommand represents a registered entity command with its configuration
type EntityCommand struct {
	Name          string         // Short name: "db", "api", "fe"
	LongName      string         // Full name: "database", "graphql", "frontend"
	EntityType    string         // Module type to filter: "database", "graphql", "frontend"
	Command       *cobra.Command // Root command for this entity
	RequireModule bool           // Whether commands require a module
}

var entityCommands = make(map[string]*EntityCommand)

// RegisterEntityCommand registers an entity command with module resolution support
func RegisterEntityCommand(config EntityCommand) {
	if _, ok := entityCommands[config.Name]; ok {
		panic(fmt.Sprintf("entity command %s is being registered twice", config.Name))
	}

	// Set up module resolution if required
	if config.RequireModule {
		setupModuleResolution(config.Command, config.EntityType)
	}

	entityCommands[config.Name] = &config
}

// GetEntityCommands returns registered entity commands ordered by short name.
func GetEntityCommands() []*EntityCommand {
	names := slices.Sorted(maps.Keys(entityCommands))
	out := make([]*EntityCommand, 0, len(names))
	for _, name := range names {
		out = append(out, entityCommands[name])
	}
	return out
}
