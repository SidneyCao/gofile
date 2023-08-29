package gofile

import (
	"errors"
	"os"
)

// Make a directory.
func (p *Path) Mkdir() error {
	if p.ifExist || p.isFile {
		return errors.New("this object already exists or is a file")
	}

	err := os.Mkdir(p.absPath, 0755)
	p.refresh(p.absPath)

	return err
}

// Make directories recursively.
func (p *Path) MkdirAll() error {
	if p.ifExist || p.isFile {
		return errors.New("this object already exists or is a file")
	}

	err := os.MkdirAll(p.absPath, 0755)
	p.refresh(p.absPath)

	return err
}

// List all files or directories in current directory.
//
// Return a slice of Path struct.
//
// Not recursively.
func (p *Path) List() ([]Path, error) {
	if p.isDir {
		paths := make([]Path, 0)
		dir, err := os.Open(p.absPath)
		if err != nil {
			return nil, err
		}
		defer dir.Close()

		subInfo, err := dir.ReadDir(0)
		if len(subInfo) > 0 {
			for _, i := range subInfo {
				p, _ := Load(p.absPath + string(os.PathSeparator) + i.Name())
				paths = append(paths, p)
			}
			return paths, err
		}
	}
	return nil, errors.New("this object is not a dir")
}
