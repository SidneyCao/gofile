package gofile

import (
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

// copy

// move

// delete

// rename
