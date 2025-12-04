package graphql

import (
	"embed"
	"io/fs"
)

//go:embed all:_template
var template embed.FS

var Template fs.FS

func init() {
	var err error
	Template, err = fs.Sub(template, "_template")
	if err != nil {
		panic(err)
	}
}
