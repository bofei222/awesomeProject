package myslices

import (
	"fmt"
	"testing"
)

var pow2 = []int{1, 2, 4, 8, 16, 32, 64, 128}

func Test27(t *testing.T) {
	for i, v := range pow2 {
		fmt.Printf("2**%d = %d\n", i, v)
	}
}
