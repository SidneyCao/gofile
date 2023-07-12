package gofile

import (
	"fmt"
	"testing"
	"time"
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
	p, _ := Load("./README.md")
	t.Run("open", func(t *testing.T) {
		p.Open()
		fmt.Println(p.file.Fd())
		if p.file == nil {
			t.Errorf("test file open error!")
		}
	})
	t.Run("close", func(t *testing.T) {
		err := p.Close()
		time.Sleep(10 * time.Second)
		fmt.Println(p.file.Fd())
		if err != nil {
			t.Errorf("test file close error!")
		}
	})
}
