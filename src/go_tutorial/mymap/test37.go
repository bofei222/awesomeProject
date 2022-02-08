package main

import "fmt"

func incr() func() int {
	var x int
	return func() int {
		x++
		return x
	}
}

func incr2() int {
	print(11)
	return 11
}

// 调用这个函数会返回一个函数变量。

func main() {
	i := incr2() // 闭包的时候 是带括号的
	print(i)
	//fmt.Println(i())
	//fmt.Println(i())
	//fmt.Println(i()) // x 逃逸了，它的生命周期没有随着它的作用域结束而结束

	fmt.Println(incr()())
	fmt.Println(incr()())
	fmt.Println(incr()())
}
