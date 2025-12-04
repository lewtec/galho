package version

import "embed"

func init() {
	var _ embed.FS // without this the go LSP will delete the import
}

//go:embed version.txt
var Version string
