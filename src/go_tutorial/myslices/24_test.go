package myslices

import (
	"fmt"
	"testing"
)

func Test24(t *testing.T) {
	var s []int
	fmt.Println(s, len(s), cap(s))
	if s == nil {
		fmt.Println("nil!")
	}
}
