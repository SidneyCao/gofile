package file

import (
	"io/fs"
	"os"
	"path/filepath"
)

type File struct {
	ifExist bool // default false
	isDir   bool // default false

	name    string // default ""
	absPath string // default ""
	ext     string // default ""

	file *fs.File // default nil
}

func Load(path string) (File, error) {
	f := File{}

	// if exist
	sts, err := os.Stat(path)
	if err != nil {
		if !os.IsExist(err) {
			f.ifExist = false
		}
		return f, err
	}
	f.ifExist = true

	// if is dir
	if sts.IsDir() {
		f.isDir = true
	}

	// get the name
	f.name = filepath.Base(path)
	// get the absPath
	f.absPath, _ = filepath.Abs(path)

	if f.isDir {
		return f, nil
	}

	// get ext
	f.ext = filepath.Ext(path)
	return f, nil
}

// for file -------------------------

// for dir --------------------------

// common ---------------------------
