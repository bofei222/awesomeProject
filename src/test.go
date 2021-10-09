package main

import (
	"fmt"
	"time"
)

func main() {
	st := time.Now()
	n := 100000000
	var j int
	z := 0
	for i := 0; i < n; i++ {
		//j += i
		//z += i
	}
	println(j)
	println(z)
	fmt.Println(fmt.Sprintf("%f seconds", float32(time.Now().Sub(st).Microseconds())/1000000))
	// 0.058280 seconds
}
