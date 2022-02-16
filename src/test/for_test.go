package test

import (
	"fmt"
	"testing"
)

func TestForChangeIndex(t *testing.T) {

	for i := 0; i < 10; i++ {
		fmt.Println(i)
		if i == 4 {
			i = 3
		}
	}
}
