package core

import (
	"fmt"
	"path/filepath"

	"github.com/h2non/findup"
)

type Project struct {
	dir string
}

// GetProject walks up from the process working directory looking for a .galho
// marker and returns the project rooted at that directory.
func GetProject() (*Project, error) {
	dotgalho, err := findup.Find(".galho")
	if err != nil {
		return nil, fmt.Errorf("not a galho project: no .galho marker found from cwd: %w", err)
	}
	// findup returns a filesystem path; use filepath (not path) so Dir is correct
	// on platforms where the separator is not '/'.
	return &Project{dir: filepath.Dir(dotgalho)}, nil
}

func (p *Project) Dir() string {
	return p.dir
}
