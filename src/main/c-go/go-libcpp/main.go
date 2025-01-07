package main

/*
#cgo CFLAGS: -I./
#cgo LDFLAGS: -L./ -lliblibcpp
#include "libcpp.h"
*/
import "C"
import (
	"fmt"
	"runtime"
)

func main() {
	// 创建并初始化 C 结构体
	data := C.Data{
		intValue:    10,
		doubleValue: 3.5,
	}

	// 调用 C++ 函数
	result := C.processData(data)

	// 输出结果
	fmt.Printf("Result: %f\n", float64(result))
	// 输出当前操作系统
	fmt.Println("当前操作系统:", runtime.GOOS)
}
