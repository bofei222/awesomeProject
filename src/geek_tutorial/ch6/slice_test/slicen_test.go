package slice_test

import (
	"fmt"
	"testing"
)

func TestForLoopRemove(t *testing.T) {
	//all := []int{0, 3, 2, 1, 4, 9, 6, 7, 8, 5}
	all := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	fmt.Println(all) //[0 1 2 3 4 5 6 7 8 9]

	selectMap := map[int]int{
		1: 0,
		3: 0,
		9: 0,
	}
	j := 0
	selectSlice := make([]int, 0)
	for _, v := range all {
		if _, ok := selectMap[v]; !ok {
			//all = RemoveIndex(all, i)
			selectSlice = append(selectSlice, v)
			j++
		}
	}
	fmt.Println(selectSlice)

}
func RemoveIndex(s []int, index int) []int {
	return append(s[:index], s[index+1:]...)
}

func TestSliceRemove(t *testing.T) {
	year := []string{"Jan", "Feb", "Mar", "Apr", "May", "Jun", "Jul", "Aug", "Sep",
		"Oct", "Nov", "Dec"}
	for i, s := range year {
		if s == "Apr" || s == "Feb" {
			year = append(year[:i], year[i+1:]...) // ... 将列表展开 为 元素
		}
	}
	fmt.Println(year)
}

func TestSliceInit(t *testing.T) {

	// 声明时 不必指定长度
	var s0 []int // 没有 ... ，所以是 切片
	s0 = append(s0, 1)

	t.Log(len(s0), cap(s0))

	s2 := make([]int, 3, 5)
	t.Log(len(s2), cap(s2))
	s3 := make([]int, 0)
	s3 = append(s3, 1)
	t.Log(s3)
	// 声明时 length ，最大可填充 cap
}

// 为什么要s= append()  地址发生了变化  。实现变长
// 切片 共享存储结构

func TestSliceShareMemory(t *testing.T) {
	year := []string{"Jan", "Feb", "Mar", "Apr", "May", "Jun", "Jul", "Aug", "Sep",
		"Oct", "Nov", "Dec"}
	Q2 := year[3:6]
	t.Log(Q2, len(Q2), cap(Q2))

	summer := year[5:8]
	t.Log(summer, len(summer), cap(summer))
	summer[0] = "UnKnow"
	t.Log(Q2) // 因为共享  受影响
	t.Log(year)

}

//数组可比较
func TestSliceComparing(t *testing.T) {

}

func TestNullArray(t *testing.T) {
	var year []string
	for _, y := range year {
		fmt.Println("y:", y)
	}

}
