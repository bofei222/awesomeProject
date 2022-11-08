package main

import "fmt"

func square(x int) int {
	//println(x*x)
	return 11
}

func square2(x int) int {
	//println(x*x)
	return 11
}
func main() {
	// 有返回值 无返回值
	i := square(1) // 非闭包的时候 带括号报错的
	//fmt.Println(i(2)) can't call non-function
	fmt.Println(i)

	j := square2
	fmt.Println(j)    //
	fmt.Println(j(2)) // 要么选择 初始化的时候 带括号，要么选择 调用的时候给参数
}
