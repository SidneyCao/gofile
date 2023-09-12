package gofile

import (
	"os"
)

type Path struct {
	IfExist bool // default false
	IsDir   bool // default false
	IsFile  bool // default false

	Name    string // default ""
	AbsPath string // default ""
	Ext     string // default ""

	File *os.File // default nil
}

// Load the path to the Path struct.
func Load(pathStr string) (Path, error) {
	p := Path{}

	err := p.refresh(pathStr)

	return p, err
}
