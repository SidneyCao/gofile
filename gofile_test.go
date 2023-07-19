package gofile

import (
	"fmt"
	"io"
	"testing"
)

// test in unix system
func TestLoad(t *testing.T) {
	t.Run("exist", func(t *testing.T) {
		p, _ := Load("./README.md")
		if !p.ifExist {
			t.Errorf("test exist error!")
		}
	})
	t.Run("not_exist", func(t *testing.T) {
		p, _ := Load("/x/y")
		if p.ifExist {
			t.Errorf("test not exist error!")
		}
	})
	t.Run("file", func(t *testing.T) {
		p, _ := Load("/etc/profile")
		if !p.isFile {
			t.Errorf("test type file error!")
		}
		if p.name != "profile" {
			t.Errorf("test file name error!")
		}
		if p.absPath != "/etc/profile" {
			t.Errorf("test file abspath error!")
		}
		p, _ = Load("./README.md")
		if p.ext != ".md" {
			t.Errorf("test file ext error!")
		}
	})
	t.Run("dir", func(t *testing.T) {
		p, _ := Load("/")
		if !p.isDir {
			t.Errorf("test type dir error!")
		}
		if p.name != "/" {
			t.Errorf("test dir name error!")
		}
		if p.absPath != "/" {
			t.Errorf("test dir abspath error!")
		}
	})
}

func TestFileOpenClose(t *testing.T) {
	p, _ := Load("./testFiles/read.txt")
	t.Run("open", func(t *testing.T) {
		err := p.Open()
		if p.file == nil {
			t.Errorf("%v, test file open error!", err)
		}
		fmt.Printf("the Path struct is %v \n", p)
	})
	t.Run("close", func(t *testing.T) {
		err := p.Close()
		if err != nil {
			t.Errorf("%v, test file close error!", err)
		}
		fmt.Printf("the file descriptor is invaild as %v \n", p.file.Fd())
	})
}

func TestFileRead(t *testing.T) {
	p, _ := Load("./testFiles/read.txt")
	t.Run("open1", func(t *testing.T) {
		err := p.Open()
		if p.file == nil {
			t.Errorf("%v, test file open error!", err)
		}
	})

	t.Run("read", func(t *testing.T) {
		b, err := p.Read()
		if err != nil && err != io.EOF {
			t.Errorf("%v, test file read error!", err)
		}
		fmt.Println(string(b))
		p.Close()
	})

	t.Run("open2", func(t *testing.T) {
		err := p.Open()
		if p.file == nil {
			t.Errorf("%v, test file open error!", err)
		}
	})

	t.Run("readline", func(t *testing.T) {
		l, err := p.ReadLines()
		if err != nil && err != io.EOF && l[2] != "3" {
			t.Errorf("%v, test file read line error!", err)
		}
		p.Close()
	})
}
