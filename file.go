package gofile

import (
	"bufio"
	"bytes"
	"errors"
	"io"
	"os"
)

// Open a file.
//
// If parent directory exist, will create if not exist.
//
// If parent directory does not exist, will throw up 'no such dir' error.
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

// Close a file.
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

// Read a file.
func (p *Path) Read() ([]byte, error) {
	if p.isDir {
		return nil, errors.New("this object is dir, can not be read")
	}

	// Move the offset to the beginning of the file
	_, err := p.file.Seek(0, 0)
	if err != nil {
		return nil, err
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

// ReadLines() will return a slice of string which contains the content of file line by line.
//
// The returned error will never be nil, most time it is io.EOF.
func (p *Path) ReadLines() ([]string, error) {
	res := make([]string, 0)
	if p.isDir {
		return res, errors.New("this object is dir, can not be read")
	}

	// Move the offset to the beginning of the file
	_, err := p.file.Seek(0, 0)
	if err != nil {
		return nil, err
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

// Receive a slice of string and write to file one by one.
//
// Write mode is append.
//
// If you want to override the file, use Truncate(0) first.
func (p *Path) Write(data []string) error {
	if p.isDir {
		return errors.New("this object is dir, can not be written")
	}
	for _, s := range data {
		n, err := p.file.WriteString(s)
		if err != nil {
			return err
		}
		if n != len(s) {
			return errors.New("write bytes num error")
		}
	}

	return nil
}

// Truncate size of a file.
func (p *Path) Truncate(size int64) error {
	if p.isDir {
		return errors.New("this object is dir, can not be truncated")
	}
	err := p.file.Truncate(size)
	return err
}
