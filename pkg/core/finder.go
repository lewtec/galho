package core

import "fmt"

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
	for k, v := range moduleFinders {
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
