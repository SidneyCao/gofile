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
	p, _ := Load("./test_files/read.txt")
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
		fmt.Printf("the file descriptor is invaild as %v, file has been close! \n", p.file.Fd())
	})
}

func TestFileRead(t *testing.T) {
	p, _ := Load("./test_files/read.txt")
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

func TestFileWrite(t *testing.T) {
	p, _ := Load("./test_files/wirte.txt")
	t.Run("open", func(t *testing.T) {
		err := p.Open()
		if p.file == nil {
			t.Errorf("%v, test file open error!", err)
		}
	})
	t.Run("truncate", func(t *testing.T) {
		err := p.Truncate(0)
		sts, _ := p.file.Stat()
		if err != nil || sts.Size() != 0 {
			t.Errorf("%v, test file truncate error!", err)
		}
	})
	t.Run("wirte", func(t *testing.T) {
		date := []string{"1\n", "22\n", "333"}
		err := p.write(date)
		if err != nil {
			t.Errorf("%v, test file write error!", err)
		}
		p.Close()
		p.Open()
		b, _ := p.Read()
		if string(b) != "1\n22\n333" {
			t.Errorf("not match,test file write error!")
		}
		p.Close()
	})
}

func TestDirMkdir(t *testing.T) {
	p, _ := Load("./test_files/test_dir")
	t.Run("mkdir", func(t *testing.T) {
		err := p.Mkdir()
		if err != nil || !p.ifExist {
			t.Errorf("%v, test dir mkdir error!", err)
		}
	})

	p, _ = Load("./test_files/test_dir/sub_dir/sub")
	t.Run("mkdirAll", func(t *testing.T) {
		err := p.MkdirAll()
		if err != nil || !p.ifExist {
			t.Errorf("%v, test dir mkdirAll error!", err)
		}
	})
}

func TestDirList(t *testing.T) {
	p, _ := Load("./")

	t.Run("list", func(t *testing.T) {
		paths, err := p.List()
		if err != nil {
			t.Errorf("%v, test dir list error!", err)
		}
		fmt.Println(paths)
	})
}
