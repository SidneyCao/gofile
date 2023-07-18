package gofile

import (
	"bufio"
	"errors"
	"io"
	"os"
)

// open file
// create if not exist
func (p *Path) Open() error {
	if p.isDir {
		return errors.New("this object is dir, can not be opened")
	}

	f, err := os.OpenFile(p.absPath, os.O_RDWR|os.O_APPEND|os.O_CREATE, 0644)
	if err != nil {
		return err
	}

	p.file = f
	p.refresh(p.absPath)
	return nil
}

// close file
func (p *Path) Close() error {
	if p.file != nil {
		err := p.file.Close()
		if err != nil {
			return err
		}
	} else {
		return errors.New("this object has not been opened, can not be closed")
	}
	return nil
}

// read file
// also will defer close file
func (p *Path) Read() ([]byte, error) {
	if p.isDir {
		return nil, errors.New("this object is dir, can not be readed")
	}
	r := bufio.NewReader(p.file)
	defer p.Close()

	res := make([]byte, 0)
	resTemp := make([]byte, 1024)

	for {
		n, err := r.Read(resTemp)
		if err != nil && err != io.EOF {
			return nil, err
		}
		if n == 0 {
			return res, err
		} else {
			res = append(res, resTemp[:n]...)
		}
	}
}

// read file by line
// also will defer close file
func (p *Path) ReadLine() (string, error) {
	if p.isDir {
		return "", errors.New("this object is dir, can not be readed")
	}
	r := bufio.NewReader(p.file)

	line, err := r.ReadSlice('\n')

	return string(line), err
}

// write file
