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
		selectedModule, err = findModuleByName(modules, moduleName)
		if err != nil {
			return nil, err
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

// findModuleByName resolves --module against discovered modules.
// Prefer a unique Module.Name() match; otherwise accept a unique path-based
// match. Multiple hits are an error so multi-module projects cannot silently
// operate on the first walker result (e.g. two modules both basename "db").
func findModuleByName(modules []Module, name string) (Module, error) {
	var exact []Module
	var partial []Module
	for _, m := range modules {
		if m.Name() == name {
			exact = append(exact, m)
			continue
		}
		if moduleMatchesName(m, name) {
			partial = append(partial, m)
		}
	}

	switch {
	case len(exact) == 1:
		return exact[0], nil
	case len(exact) > 1:
		return nil, ambiguousModuleError(name, exact)
	case len(partial) == 1:
		return partial[0], nil
	case len(partial) > 1:
		return nil, ambiguousModuleError(name, partial)
	default:
		return nil, fmt.Errorf("module %s not found", name)
	}
}

func ambiguousModuleError(name string, matches []Module) error {
	paths := make([]string, 0, len(matches))
	for _, m := range matches {
		paths = append(paths, m.Path())
	}
	return fmt.Errorf("module %q is ambiguous; matches: %s", name, strings.Join(paths, ", "))
}

func moduleMatchesName(m Module, name string) bool {
	// Prefer the friendly module name used in listings and task ids
	// (e.g. DatabaseModule names "crm" from internal/crm/db).
	if m.Name() == name {
		return true
	}

	// Path-based fallbacks: "internal/crm/db" matches "crm" or "db".
	path := m.Path()

	parentDir := filepath.Base(filepath.Dir(path))
	if parentDir == name {
		return true
	}

	dirName := filepath.Base(path)
	if dirName == name {
		return true
	}

	parts := strings.Split(filepath.ToSlash(path), "/")
	for _, part := range parts {
		if part == name {
			return true
		}
	}

	return false
}
