package main

/*
#include <stdio.h>

int add(int a, int b) {
    return a + b;
}
*/
import "C"
import "fmt"

func main() {
	// 使用 C 函数
	a, b := 3, 4
	result := C.add(C.int(a), C.int(b)) // 调用 C 函数
	fmt.Println("Result of addition:", result)
}
