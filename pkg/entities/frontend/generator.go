package frontend

import (
	"github.com/lewtec/galho/pkg/core"
	"github.com/lewtec/galho/pkg/utils/scaffold"
)

func init() {
	core.RegisterGenerateCommand("frontend", core.NewScaffoldGenerateCommand(
		"frontend [path]",
		"Generate a frontend module",
		func(path string) error { return scaffold.InstallFS(path, Template) },
	))
}
