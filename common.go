package gofile

import (
	"errors"
	"os"
	"path/filepath"
)

// refresh the Path struct info
func (p *Path) refresh(pathStr string) error {
	// get the name
	p.name = filepath.Base(pathStr)
	// get the absPath
	p.absPath, _ = filepath.Abs(pathStr)

	// if exist
	sts, err := os.Stat(pathStr)
	if err != nil {
		if !os.IsExist(err) {
			p.ifExist = false
		}
		return err
	}
	p.ifExist = true

	// if is dir
	if sts.IsDir() {
		p.isDir = true
		return nil
	} else {
		p.isFile = true
	}

	// get ext
	p.ext = filepath.Ext(pathStr)
	return nil
}

// rename
//func (p *Path) Rename() error {
//}

// delete a file
// or delete all files of a dir
func (p *Path) Delete() error {
	if !p.ifExist {
		return errors.New("this object does not exist, can not be delete")
	} else {
		if p.isFile {
			err := os.Remove(p.absPath)
			if err != nil {
				return err
			}
			p.refresh(p.absPath)
		} else {
			err := os.RemoveAll(p.absPath)
			if err != nil {
				return err
			}
			p.refresh(p.absPath)
		}
	}
	return nil
}

// move

// copy
