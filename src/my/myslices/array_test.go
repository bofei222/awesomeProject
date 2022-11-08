package myslices

import (
	"fmt"
	"testing"
)

func TestArray(t *testing.T) {
	var a [2]string
	a[0] = "Hello"
	a[1] = "World"
	fmt.Println(a[0], a[1])

	a[1] = "World1"
	fmt.Println(a)
}
