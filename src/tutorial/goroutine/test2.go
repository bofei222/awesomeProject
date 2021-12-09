package main

import (
	"fmt"
	"time"
)

func main() {
	for i := 0; i < 10; i++ {
		var name = i
		go func() {
			fmt.Println(name)
		}()
	}
	time.Sleep(2 * time.Second)
}
