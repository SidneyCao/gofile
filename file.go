package gofile

import (
	"bufio"
	"bytes"
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
func (p *Path) Read() ([]byte, error) {
	if p.isDir {
		return nil, errors.New("this object is dir, can not be readed")
	}
	r := bufio.NewReader(p.file)

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

// ReadLine will return a slice of string
// which contains the content of file line by line
// the return error will never be nil, most time it is io.EOF
func (p *Path) ReadLine() ([]string, error) {
	res := make([]string, 0)
	if p.isDir {
		return res, errors.New("this object is dir, can not be readed")
	}

	r := bufio.NewReader(p.file)

	for {
		line, err := r.ReadBytes('\n')
		line = bytes.TrimRight(line, "\r\n")
		res = append(res, string(line))
		if err != nil {
			// if err == io.EOF will also return
			return res, err
		}
	}
}

// write file
