package main

import (
	"awesomeProject/src/another"
	_ "unsafe"
)

//go:linkname m somewhere.com/someone/another.m
var m map[int]string

func main() {
	println(m[1])
	//m[2] = "b"
	println(another.M(2))
}
