# gofile
Package gofile provide some utilities for working with local files in golang.  
<br>
[![go.dev](https://img.shields.io/badge/go.dev-reference-007d9c?logo=go&logoColor=white&style=flat-square)](https://pkg.go.dev/github.com/sidneycao/gofile?tab=doc)
[![codecov](https://codecov.io/gh/sidneycao/gofile/graph/badge.svg?token=YJVHRZAL1X)](https://codecov.io/gh/sidneycao/gofile)
[![Go Report Card](https://goreportcard.com/badge/github.com/sidneycao/gofile)](https://goreportcard.com/report/github.com/sidneycao/gofile)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)
[![FOSSA Status](https://app.fossa.com/api/projects/git%2Bgithub.com%2Fsidneycao%2Fgofile.svg?type=shield&issueType=license)](https://app.fossa.com/projects/git%2Bgithub.com%2Fsidneycao%2Fgofile?ref=badge_shield)  

## Install  
```sh
  $ go get github.com/sidneycao/gofile
``` 
  
## Example:  
Here is an example of working with local files, more functions can be found in [Documentation](https://pkg.go.dev/github.com/sidneycao/gofile#section-documentation).  
```go
package main

import (
	"fmt"
	"io"
	"log"

	"github.com/sidneycao/gofile"
)

func main() {
	// Load a local file.
	p, err := gofile.Load("./example-file.txt")
	if err != nil {
		// If the file does not exist, will throw up 'no such file' error.
		//
		// But if the parent exist, this file will be create by p.Open().
		log.Println(err)
	}

	// Open a file.
	err = p.Open()
	if err != nil {
		log.Panic(err)
	}

	// Write to a file.
	wData := []string{"hello\n", "i am\n", "a buff\n"}
	err = p.Write(wData)
	if err != nil {
		log.Panic(err)
	}

	// Read a file.
	rData, err := p.Read()
	if err != nil && err != io.EOF {
		log.Println(err)
	}
	fmt.Println(string(rData))

	// Truncate a file.
	err = p.Truncate(0)
	if err != nil {
		log.Panic(err)
	}

	// Read a file line by line.
	wDataLine := []string{"line1\n", "line2\n", "line3\n"}
	err = p.Write(wDataLine)
	if err != nil {
		log.Panic(err)
	}
	rDataLine, err := p.ReadLines()
	if err != nil && err != io.EOF {
		log.Panic(err)
	}
	fmt.Println(rDataLine)

	// remeber to close.
	p.Close()
}
```  

