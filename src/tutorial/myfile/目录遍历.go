package main

import (
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
)

func main() {
	file, err := os.OpenFile("D:/home", os.O_RDWR, os.ModeDir)
	if nil != err {
		fmt.Println(err)
	}
	fmt.Println(file.Stat())

	filepath.Walk("D:/home", func(path string, info fs.FileInfo, err error) error {
		fmt.Println(path)
		return err
	})

}
