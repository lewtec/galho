package core

import "path/filepath"

// BaseModule provides common implementation for Module interface
type BaseModule struct {
	path string
	name string
}

// NewBaseModule creates a new BaseModule
func NewBaseModule(path string, name string) BaseModule {
	return BaseModule{
		path: path,
		name: name,
	}
}

// Path returns the path to the module relative to the project root
func (m *BaseModule) Path() string {
	return m.path
}

// Name returns a friendly name for the module
func (m *BaseModule) Name() string {
	return m.name
}

// DeriveModuleName derives a module name from its path.
// It uses the parent directory name as the module name.
func DeriveModuleName(path string) string {
	return filepath.Base(filepath.Dir(path))
}
