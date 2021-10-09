package main

import (
	"fmt"
	"math"
)

func Sqrt(x float64) float64 {
	root := 100.0
	for math.Abs((root*root - x)) > 0.00000000001 {
		root = (root + x/root) / 2
	}
	return root
}

func main() {
	fmt.Println(math.Sqrt(2))
	fmt.Println(Sqrt(2))
}
