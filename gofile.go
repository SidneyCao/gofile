package gofile

import (
	"os"
	"path/filepath"
)

type Path struct {
	ifExist bool // default false
	isDir   bool // default false
	isFile  bool // default false

	name    string // default ""
	absPath string // default ""
	ext     string // default ""

	file *os.File // default nil
}

func Load(pathStr string) (Path, error) {
	p := Path{}

	// if exist
	sts, err := os.Stat(pathStr)
	if err != nil {
		if !os.IsExist(err) {
			p.ifExist = false
		}
		return p, err
	}
	p.ifExist = true

	// if is dir
	if sts.IsDir() {
		p.isDir = true
	} else {
		p.isFile = true
	}

	// get the name
	p.name = filepath.Base(pathStr)
	// get the absPath
	p.absPath, _ = filepath.Abs(pathStr)

	if p.isDir {
		return p, nil
	}

	// get ext
	p.ext = filepath.Ext(pathStr)
	return p, nil
}
