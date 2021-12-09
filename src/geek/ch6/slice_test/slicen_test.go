package slice_test

import (
	"fmt"
	"testing"
)

func TestSliceInit(t *testing.T) {

	// 声明时 不必指定长度
	var s0 []int // 没有 ... ，所以是 切片
	s0 = append(s0, 1)

	t.Log(len(s0), cap(s0))

	s2 := make([]int, 3, 5)
	t.Log(len(s2), cap(s2))
	s2 = append(s2, 1)
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
