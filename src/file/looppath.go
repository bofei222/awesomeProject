package main

import (
	"fmt"
	"io/ioutil"
)

func main() {
	GetAllFile("D:\\home")
}

func GetAllFile(pathname string) error {
	rd, err := ioutil.ReadDir(pathname)
	for _, fi := range rd {
		if fi.IsDir() {
			fmt.Printf("[%s]\n", pathname+"\\"+fi.Name())
			GetAllFile(pathname + fi.Name() + "\\")

		} else {
			fmt.Println(fi.Name())
			//filepath.Abs(filepath.Dir(fi))
			//open, _ := os.Open(fi.Name())
			//fmt.Println(open.Name())
		}
	}
	return err
}
