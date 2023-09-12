package gofile

import (
	"errors"
	"os"
	"path/filepath"
)

// Refresh the Path struct info.
func (p *Path) refresh(pathStr string) error {
	// get the name
	p.Name = filepath.Base(pathStr)
	// get the absPath
	p.AbsPath, _ = filepath.Abs(pathStr)

	// if exist
	sts, err := os.Stat(pathStr)
	if err != nil {
		if !os.IsExist(err) {
			p.IfExist = false
		}
		return err
	}
	p.IfExist = true

	// if is dir
	if sts.IsDir() {
		p.IsDir = true
		return nil
	} else {
		p.IsFile = true
	}

	// get ext
	p.Ext = filepath.Ext(pathStr)
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
	if !p.IfExist {
		return errors.New("this object does not exist, can not be moved")
	}
	if newPath[len(newPath)-1] == os.PathSeparator {
		newPath = filepath.Join(filepath.Dir(newPath), p.Name)
	}

	err := os.Rename(p.AbsPath, newPath)
	if err != nil {
		return err
	}

	p.refresh(newPath)
	return nil
}

// Delete a file or delete a directory with its all files.
func (p *Path) Delete() error {
	if !p.IfExist {
		return errors.New("this object does not exist, can not be deleted")
	}

	if p.IsFile {
		err := os.Remove(p.AbsPath)
		if err != nil {
			return err
		}
	} else {
		err := os.RemoveAll(p.AbsPath)
		if err != nil {
			return err
		}
	}

	p.refresh(p.AbsPath)
	return nil
}
