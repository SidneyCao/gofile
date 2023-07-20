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

	err := os.Mkdir(p.absPath, 0755)
	p.refresh(p.absPath)

	return err
}

// mkdir recursively
func (p *Path) mkdirAll() error {
	if p.ifExist || p.isFile {
		return errors.New("this object already exists or is a file")
	}

	err := os.MkdirAll(p.absPath, 0755)
	p.refresh(p.absPath)

	return err
}

// list
func (p *Path) list() {

}

//
