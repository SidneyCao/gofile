package gofile

import (
	"errors"
	"os"
)

// open file
func (p *Path) Open() error {
	if p.isDir {
		return errors.New("this object is dir, can not be opened")
	}

	f, err := os.OpenFile(p.absPath, os.O_RDWR|os.O_APPEND|os.O_CREATE, 0644)
	if err != nil {
		return err
	}

	p.file = f

	return nil
}

// close file
func (p *Path) Close() error {
	if p.file != nil {
		err := p.file.Close()
		if err != nil {
			return err
		}
	} else {
		return errors.New("this object has not been opened, can not be closed")
	}
	return nil
}

// create file
func (p *Path) Create() error {
	if p.ifExist {
		return errors.New("this object is existed, can not be created")
	}
	f, err := os.Create(p.absPath)
	if err != nil {
		return err
	}
	p.ifExist = true
	p.isFile = true
	p.file = f
	return nil
}
