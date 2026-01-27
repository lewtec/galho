package core

import (
	"fmt"

	"github.com/spf13/cobra"
)

const moduleFlagName = "module"

// setupModuleResolution adds module resolution to an entity command
func setupModuleResolution(cmd *cobra.Command, entityType string) {
	// Add --module flag to the root command
	cmd.PersistentFlags().StringP(moduleFlagName, "m", "",
		fmt.Sprintf("Specify the %s module to operate on", entityType))

	// Add persistent pre-run hook for module resolution
	originalPreRunE := cmd.PersistentPreRunE
	cmd.PersistentPreRunE = func(cmd *cobra.Command, args []string) error {
		// Run original pre-run if exists
		if originalPreRunE != nil {
			if err := originalPreRunE(cmd, args); err != nil {
				return err
			}
		}

		// Resolve module
		ctx, err := resolveModule(cmd, entityType)
		if err != nil {
			return err
		}

		// Store context
		SetCommandContext(cmd, ctx)
		return nil
	}
}

// resolveModule finds and selects the appropriate module
func resolveModule(cmd *cobra.Command, entityType string) (*CommandContext, error) {
	// Get project
	project, err := GetProject()
	if err != nil {
		return nil, fmt.Errorf("failed to get project: %w", err)
	}

	// Check for --module flag
	moduleName, _ := cmd.Flags().GetString(moduleFlagName)

	// Find all modules of this type
	var modules []Module
	err = project.FindModules(func(found ModuleFound) bool {
		if found.Module.Type() == entityType {
			modules = append(modules, found.Module)
		}
		return true
	})
	if err != nil {
		return nil, fmt.Errorf("failed to find modules: %w", err)
	}

	// Select module using strategy
	selector := &StandardSelector{}
	selectedModule, err := selector.Select(modules, entityType, moduleName)
	if err != nil {
		return nil, err
	}

	return &CommandContext{
		Project: project,
		Module:  selectedModule,
	}, nil
}
