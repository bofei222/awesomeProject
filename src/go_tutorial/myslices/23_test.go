package myslices

import (
	"fmt"
	"testing"
)

func remove(s []int, i int) []int {
	s[i] = s[len(s)-1]
	return s[:len(s)-1]
}

func Test233(t *testing.T) {
	s := []int{2, 3, 5, 7, 11, 13}
	printSlice(s)

	ints := remove(s, 0)
	printSlice(ints)
}

func Test23(t *testing.T) {
	s := []int{2, 3, 5, 7, 11, 13}
	printSlice(s)

	// 截取切片使其长度为 0
	s = s[:0]
	printSlice(s)

	// 拓展其长度
	s = s[:4]
	printSlice(s)

	// 舍弃前两个值
	s = s[2:]
	printSlice(s)
}

func printSlice(s []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}
