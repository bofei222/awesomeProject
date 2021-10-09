package main

import (
	"fmt"
	"os"
)

func main() {
	path := "D:/home2223"
	_, err := os.Stat(path)
	os.IsNotExist(err)

	if _, err := os.Stat(path); os.IsNotExist(err) {
		err := os.Mkdir(path, os.ModePerm)
		// TODO: handle error
		fmt.Println(err)
	}
}
