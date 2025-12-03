package core

import (
	"fmt"
	"path/filepath"
	"strings"

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

	if len(modules) == 0 {
		return nil, fmt.Errorf("no %s modules found in project", entityType)
	}

	// Select module
	var selectedModule Module
	if moduleName != "" {
		// Find module by name
		selectedModule = findModuleByName(modules, moduleName)
		if selectedModule == nil {
			return nil, fmt.Errorf("module %s not found", moduleName)
		}
	} else if len(modules) == 1 {
		// Auto-select if only one module
		selectedModule = modules[0]
		fmt.Printf("Auto-selected module: %s\n", selectedModule.Path())
	} else {
		// Interactive selection
		selectedModule, err = selectModuleInteractive(modules, entityType)
		if err != nil {
			return nil, err
		}
	}

	return &CommandContext{
		Project: project,
		Module:  selectedModule,
	}, nil
}

func findModuleByName(modules []Module, name string) Module {
	for _, m := range modules {
		// Match against path components
		if moduleMatchesName(m, name) {
			return m
		}
	}
	return nil
}

func moduleMatchesName(m Module, name string) bool {
	// Implementation: check if module path contains name
	// e.g., "internal/crm/db" matches "crm"
	path := m.Path()

	// Try matching the parent directory name (e.g., "crm" in "internal/crm/db")
	parentDir := filepath.Base(filepath.Dir(path))
	if parentDir == name {
		return true
	}

	// Try matching the directory name itself
	dirName := filepath.Base(path)
	if dirName == name {
		return true
	}

	// Try matching against any path component
	parts := strings.Split(filepath.ToSlash(path), "/")
	for _, part := range parts {
		if part == name {
			return true
		}
	}

	return false
}
