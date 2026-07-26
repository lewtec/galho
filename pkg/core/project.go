package core

import (
	"errors"
	"fmt"
	"path/filepath"

	"github.com/h2non/findup"
)

// ErrNotAGalhoProject is returned when GetProject cannot find a .galho marker
// walking up from the process working directory.
var ErrNotAGalhoProject = errors.New("not a galho project")

type Project struct {
	dir string
}

// GetProject walks up from the process working directory looking for a .galho
// marker and returns the project rooted at that directory.
func GetProject() (*Project, error) {
	dotgalho, err := findup.Find(".galho")
	if err != nil {
		return nil, fmt.Errorf("%w: no .galho marker found from cwd: %w", ErrNotAGalhoProject, err)
	}
	// findup returns a filesystem path; use filepath (not path) so Dir is correct
	// on platforms where the separator is not '/'.
	return &Project{dir: filepath.Dir(dotgalho)}, nil
}

func (p *Project) Dir() string {
	return p.dir
}
