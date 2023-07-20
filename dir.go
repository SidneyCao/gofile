package gofile

import (
	"errors"
	"os"
)

// mkdir
func (p *Path) mkdir() error {
	if p.ifExist || p.isFile {
		return errors.New("this object already exists or is a file")
	}

	err := os.Mkdir(p.absPath, 0644)
	if err != nil {
		return err
	}

	p.refresh(p.absPath)
	return nil
}

// mkdir recursively

// list

//
