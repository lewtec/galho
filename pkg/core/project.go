package core

import (
	"errors"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
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
	dir, err := os.Getwd()
	if err != nil {
		return nil, fmt.Errorf("%w: get working directory: %w", ErrNotAGalhoProject, err)
	}
	start := dir
	for {
		marker := filepath.Join(dir, ".galho")
		if _, err := os.Stat(marker); err == nil {
			return &Project{dir: dir}, nil
		} else if !errors.Is(err, fs.ErrNotExist) {
			return nil, fmt.Errorf("%w: stat %s: %w", ErrNotAGalhoProject, marker, err)
		}
		parent := filepath.Dir(dir)
		if parent == dir {
			return nil, fmt.Errorf("%w: no .galho marker found from cwd %s", ErrNotAGalhoProject, start)
		}
		dir = parent
	}
}

func (p *Project) Dir() string {
	return p.dir
}
