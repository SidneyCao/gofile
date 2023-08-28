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

// p.Move("/example/"): move to another location under the dir named example.
//
// p.Move("/example"): move to another location named example
// move under the same dir is equivalent to rename
// if parent dir does not exist, will throw no such dir error
func (p *Path) Move(newPath string) error {
	if !p.ifExist {
		return errors.New("this object does not exist, can not be moved")
	}

	// get the abs path
	// newAbsPath := ""
	/*
		if filepath.IsAbs(newPath) {
			if newPath[len(newPath)-1] == os.PathSeparator {
				newAbsPath = filepath.Join(newPath, p.name)
			} else {
				newAbsPath = newPath
			}
		} else {
			if newPath[len(newPath)-1] == os.PathSeparator {
				newAbsPath = filepath.Join(filepath.Dir(newPath), p.name)
			} else {
				newAbsPath = filepath.Dir(newPath)
			}
		}
	*/
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

// delete a file
// or delete all files of a dir
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

// copy
