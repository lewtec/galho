package scaffold

import "io/fs"

// MustSub returns the named subdirectory of fsys, or panics.
// Entity packages use this for their //go:embed all:_template roots so the
// Sub + panic-on-error boilerplate is not copied into every entity package.
func MustSub(fsys fs.FS, dir string) fs.FS {
	sub, err := fs.Sub(fsys, dir)
	if err != nil {
		panic(err)
	}
	return sub
}
