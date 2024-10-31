package main

import (
	"fmt"
	"unsafe"
)

// 定义一个结构体，包含一些大字段
type LargeStruct struct {
	Numbers [1000]float64 // 假设结构体包含 1000 个浮点数
}

func main() {
	// 创建 1000 个 LargeStruct 实例，不使用指针
	nonPointerSlice := make([]LargeStruct, 1000)
	fmt.Printf("Non-pointer slice memory usage: %d bytes\n", unsafe.Sizeof(nonPointerSlice[0])*1000)

	// 创建 1000 个指向 LargeStruct 的指针
	pointerSlice := make([]*LargeStruct, 1000)
	for i := range pointerSlice {
		pointerSlice[i] = &LargeStruct{}
	}
	fmt.Printf("Pointer slice memory usage: %d bytes\n", unsafe.Sizeof(pointerSlice[0])*1000)
}
