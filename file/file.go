package file

import "os"

type File struct {
	isExist bool
	isFile  bool

	name    string
	absPath string
	ext     string

	file *os.File
}
