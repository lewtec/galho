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

func (p *Project) Dir() string {
	return p.dir
}
