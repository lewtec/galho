package database

import (
	"embed"
	"io/fs"
)

//go:embed all:template
var template embed.FS

var Template fs.FS

func init() {
	var err error
	Template, err = fs.Sub(template, "template")
	if err != nil {
		panic(err)
	}
}
