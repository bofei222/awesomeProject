package main

import "fmt"

func main() {
	b := 6

	var b_ptr *int // *int is used to declare variable
	// b_ptr to be a pointer to an int

	b_ptr = &b // b_ptr is assigned the value that is the
	// address of where variable b is stored

	// Shorthand for the above two lines is:
	// b_ptr := &b

	fmt.Printf("address of b_ptr: %p\n", b_ptr) // 打印 地址

	// We can use *b_ptr to get the value that is stored
	// at address b_ptr, known as dereferencing the pointer
	fmt.Printf("value stored at b_ptr: %d\n", *b_ptr) // 打印指针指向值。

}
