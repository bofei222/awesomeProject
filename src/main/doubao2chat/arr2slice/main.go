package main

import "fmt"

func modifyArray(arr [5]int) {
	arr[0] = 100
}

func modifySlice(slice []int) {
	slice[0] = 100
}

func main() {
	arr := [5]int{1, 2, 3, 4, 5}
	slice := []int{1, 2, 3, 4, 5}

	modifyArray(arr) // arr 不会改变
	fmt.Println(arr)
	modifySlice(slice) // slice 会改变
	fmt.Println(slice)
}
