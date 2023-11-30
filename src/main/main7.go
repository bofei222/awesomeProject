package main

import (
	"fmt"
)

func main() {
	var a string = "hello world"
	fmt.Println(a)     // output: hello world
	fmt.Println(&a)    // output: 0xc00010a220 (won’t be this exactly)
	var b *string = &a //declare “b” as type “pointer to string”
	fmt.Println(b)     // output: 0xc00010a220 (wont’ be this exactly)
	fmt.Println(*b)    // output: hello world
}

func main2() {
	var a string = "hello world"
	fmt.Println(a)  // output: hello world
	fmt.Println(&a) // output: 0xc00010a220 (won’t be this exactly)
}
