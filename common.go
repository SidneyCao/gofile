package gofile

import (
	"errors"
	"os"
	"path/filepath"
)

// Name
func (p *Path) Name() string {
	return p.name
}

// is Dir
func (p *Path) IsDir() bool {
	return p.isDir
}

// If exist
func (p *Path) IfExist() bool {
	return p.ifExist
}

// Refresh the Path struct info.
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

// p.Move("/example/"): Move to another location under the directory named 'example'.
//
// p.Move("/example"): Move to another location named 'example'.
//
// Move under the same directory is equivalent to rename.
//
// If parent directory does not exist, will throw up 'no such dir' error.
func (p *Path) Move(newPath string) error {
	if !p.ifExist {
		return errors.New("this object does not exist, can not be moved")
	}
	if newPath[len(newPath)-1] == os.PathSeparator {
		newPath = filepath.Join(filepath.Dir(newPath), p.name)
	}

	err := os.Rename(p.absPath, newPath)
	if err != nil {
		return err
	}

	p.refresh(newPath)
	return nil
}

// Delete a file or delete a directory with its all files.
func (p *Path) Delete() error {
	if !p.ifExist {
		return errors.New("this object does not exist, can not be deleted")
	}

	if p.isFile {
		err := os.Remove(p.absPath)
		if err != nil {
			return err
		}
	} else {
		err := os.RemoveAll(p.absPath)
		if err != nil {
			return err
		}
	}

	p.refresh(p.absPath)
	return nil
}
