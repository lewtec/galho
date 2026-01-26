package core

import (
	"path"

	"github.com/h2non/findup"
)

type Project struct {
	dir string
}

func GetProject() (*Project, error) {
	dotgalho, err := findup.Find(".galho")
	if err != nil {
		return nil, err
	}
	projectRoot := path.Dir(dotgalho)
	return &Project{projectRoot}, nil
}

// NewProject creates a new project instance with the given directory.
// This is primarily used for testing or when the project directory is already known.
func NewProject(dir string) *Project {
	return &Project{dir: dir}
}

func (p *Project) Dir() string {
	return p.dir
}
