package main

/*
#cgo CFLAGS: -I${SRCDIR}
#cgo LDFLAGS: -L${SRCDIR} -lliblibc

#include "library.h"
*/
import "C"
import "fmt"

func main() {
	a, b := 10, 5

	// 调用 C 函数 add
	sum := C.add(C.int(a), C.int(b))
	fmt.Printf("Result of %d + %d = %d\n", a, b, int(sum))

	// 调用 C 函数 subtract
	diff := C.subtract(C.int(a), C.int(b))
	fmt.Printf("Result of %d - %d = %d\n", a, b, int(diff))
}
