package main

import "fmt"

type A struct {
	Name string
}

func (a A) foo() { //接收者写在函数名前面的括号里面
	fmt.Println("foo")
}

func main() {
	a := A{}
	a.foo() //foo
}
