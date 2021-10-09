package main

import (
	"fmt"
	"os"
)

func main() {
	var path string
	fmt.Scan(&path)

	//
	file, err := os.OpenFile(path, os.O_RDONLY, os.ModeDir)
	if err != nil {

	}
	defer file.Close()
	files, err := file.ReadDir(-1)

	for _, fileInfo := range files {
		fmt.Println(fileInfo.Name())
	}

}
