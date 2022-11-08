package mytest

import (
	"fmt"
	"math"
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

func TestType(t *testing.T) {
	var a float64 = -1
	var b float64 = 0
	c := a / b
	if math.IsInf(c, -1) {
		fmt.Print(c)
	}

}
