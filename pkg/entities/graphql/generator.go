package graphql

import (
	"github.com/lewtec/galho/pkg/core"
	"github.com/lewtec/galho/pkg/utils/scaffold"
)

func init() {
	core.RegisterGenerateCommand("graphql", core.NewScaffoldGenerateCommand(
		"graphql [path]",
		"Generate a graphql module",
		func(path string) error { return scaffold.InstallFS(path, Template) },
	))
}
