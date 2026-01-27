package core

import (
	"fmt"
	"path/filepath"
	"strings"
)

// ModuleSelector handles the logic for selecting a module from a list.
type ModuleSelector interface {
	Select(modules []Module, entityType string, preferredName string) (Module, error)
}

// StandardSelector is the default implementation of ModuleSelector.
type StandardSelector struct{}

// Select selects a module based on the preferred name, or interactively if needed.
func (s *StandardSelector) Select(modules []Module, entityType string, preferredName string) (Module, error) {
	if len(modules) == 0 {
		return nil, fmt.Errorf("no %s modules found in project", entityType)
	}

	if preferredName != "" {
		// Find module by name
		selectedModule := s.findModuleByName(modules, preferredName)
		if selectedModule == nil {
			return nil, fmt.Errorf("module %s not found", preferredName)
		}
		return selectedModule, nil
	}

	if len(modules) == 1 {
		// Auto-select if only one module
		selectedModule := modules[0]
		fmt.Printf("Auto-selected module: %s\n", selectedModule.Path())
		return selectedModule, nil
	}

	// Interactive selection
	return selectModuleInteractive(modules, entityType)
}

func (s *StandardSelector) findModuleByName(modules []Module, name string) Module {
	for _, m := range modules {
		// Match against path components
		if s.moduleMatchesName(m, name) {
			return m
		}
	}
	return nil
}

func (s *StandardSelector) moduleMatchesName(m Module, name string) bool {
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
