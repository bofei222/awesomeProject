package main

import (
	"fmt"
	"os"
	"sort"
)

func main() {
	dir, _ := ReadDir("D:/home")
	for i := range dir {
		fmt.Println(dir[i])

	}
}

func ReadDir(dirname string) ([]os.FileInfo, error) {
	f, err := os.Open(dirname)
	if err != nil {

		return nil, err
	}
	list, err := f.Readdir(-1)
	f.Close()
	if err != nil {
		return nil, err
	}
	sort.Slice(list, func(i, j int) bool { return list[i].Name() < list[j].Name() })
	return list, nil
}
