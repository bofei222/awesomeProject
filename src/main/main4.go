package main

import "fmt"

type Event struct {
	MessageId string
}

func ModifySlice(s []int) {
	s[0] = 999

}

func ModifySlice3(s []Event) {
	s[0].MessageId = "999"
}

func ModifySlice4(events []Event) {
	for i, event := range events {
		event.MessageId = "995"
		events[i] = event
	}

}
func main() {
	numbers := []int{1, 2, 3, 4, 5}

	// 调用函数，传递切片
	ModifySlice(numbers)

	// 原始切片已被修改
	fmt.Println(numbers) // 输出：[999 2 3 4 5]

	ModifySlice2(numbers...)
	fmt.Println(numbers) // 输出：[998 2 3 4 5]

	// 传递切片的引用
	events := []Event{{"1"}, {"2"}, {"3"}}
	//
	ModifySlice3(events)
	fmt.Println(events) // 输出：[{999} {2} {3}]

	ModifySlice4(events)
	fmt.Println(events) // 输出：[{995} {995} {995}]

}

func ModifySlice2(s ...int) {
	s[0] = 998
}
