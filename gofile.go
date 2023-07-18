package gofile

import (
	"os"
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

	err := p.refresh(pathStr)

	return p, err
}
