package database

import (
	"github.com/lewtec/galho/pkg/core"
	"github.com/lewtec/galho/pkg/utils/scaffold"
)

func init() {
	core.RegisterGenerateCommand("database", core.NewScaffoldGenerateCommand(
		"database [path]",
		"Generate a database module",
		func(path string) error { return scaffold.InstallFS(path, Template) },
	))
}
