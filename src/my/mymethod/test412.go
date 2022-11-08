package main

import (
	"fmt"
)

type myInt int

func Add(a, b int) { //函数
	fmt.Println(a + b)
}

func (a myInt) Add(b int) { //方法
	fmt.Println(a)
	fmt.Println(b)
}

func main() {
	a, b := 3, 4

	Add(a, b)

	var aa myInt = 3
	var bb = 4
	aa.Add(bb)

}
