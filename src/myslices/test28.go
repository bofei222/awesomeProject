package myslices

import "fmt"

func main() {
	pow := make([]int, 10)

	for i := range pow {
		pow[i] = 1 << uint(i) // == 2**i
		println(pow[i])
	}
	println("11111111111")
	for _, value := range pow {
		fmt.Printf("%d\n", value)
	}
}
