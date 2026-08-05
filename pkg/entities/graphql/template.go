package graphql

import (
	"embed"
	"io/fs"

	"github.com/lewtec/galho/pkg/utils/scaffold"
)

//go:embed all:_template
var template embed.FS

// Template is the entity's scaffold tree (contents of _template/).
var Template fs.FS = scaffold.MustSub(template, "_template")
