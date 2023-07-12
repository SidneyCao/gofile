package gofile

import (
	"errors"
	"os"
)

func (p *Path) Open() error {
	if p.isDir {
		return errors.New("this object is dir, can not be opened")
	}

	f, err := os.OpenFile(p.absPath, os.O_RDWR|os.O_APPEND, 0755)
	if err != nil {
		return err
	}

	p.file = f

	return nil
}

func (p *Path) Close() error {
	if p.file != nil {
		err := p.file.Close()
		if err != nil {
			return err
		}
	}
	return nil
}
