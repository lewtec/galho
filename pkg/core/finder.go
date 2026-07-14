package core

import (
	"fmt"
	"maps"
	"os"
	"path/filepath"
	"slices"
	"strings"
)

var moduleFinders = make(map[string]func(*Project) ([]Module, error))

type ModuleFound struct {
	Finder string
	Module Module
}

func RegisterModuleFinder(name string, t func(*Project) ([]Module, error)) {
	if _, ok := moduleFinders[name]; ok {
		panic(fmt.Sprintf("resource type %s is being registered twice", name))
	}
	moduleFinders[name] = t
}

func (p *Project) FindModules(yield func(ModuleFound) bool) error {
	// Stable order so module listing and task generation are deterministic.
	for _, k := range slices.Sorted(maps.Keys(moduleFinders)) {
		v := moduleFinders[k]
		modules, err := v(p)
		if err != nil {
			return fmt.Errorf("on finding modules with finder %s: %w", k, err)
		}
		for _, module := range modules {
			found := ModuleFound{
				Module: module,
				Finder: k,
			}
			if !yield(found) {
				return nil
			}
		}
	}
	return nil
}

// WalkModules walks the project directory and calls matchFunc on each file.
// It skips common directories like .git and node_modules.
func WalkModules(p *Project, matchFunc func(path string, info os.FileInfo) (Module, error)) ([]Module, error) {
	var modules []Module

	err := filepath.Walk(p.Dir(), func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if info.IsDir() {
			// Skip .git, node_modules, etc
			if strings.HasPrefix(info.Name(), ".") || info.Name() == "node_modules" {
				return filepath.SkipDir
			}
			return nil
		}

		mod, err := matchFunc(path, info)
		if err != nil {
			return err
		}
		if mod != nil {
			modules = append(modules, mod)
		}

		return nil
	})

	return modules, err
}
