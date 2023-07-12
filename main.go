package main

import (
	"fmt"
	"gio/file"
	"log"
)

func main() {
	f, err := file.Load("./")
	if err != nil {
		log.Println(err)
	}
	fmt.Println(f)
}
