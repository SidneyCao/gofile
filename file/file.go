package file

import (
	"io/fs"
)

type File struct {
	isExist bool
	isFile  bool

	name    string
	absPath string
	ext     string

	file *fs.File
}

func Load(path string) (File, error) {
	f := File{}

	// check the path validation
	// invalid types as following
	// " /x", "../x", "./x", "/./x", "/../x"
	if !fs.ValidPath(path) {
		return f, &fs.PathError{
			Op:   "open",
			Path: path,
			Err:  fs.ErrInvalid,
		}
	}

	return f, nil
}

// for file -------------------------

// for dir --------------------------

// common
