package arrat_test

import "testing"

func TestArrayTravel(t *testing.T) {
	arr3 := [...]int{1, 3, 4, 5}
	for index, e := range arr3 {
		t.Log(index, e)
	}

}

func TestArraySection(t *testing.T) {
	arr3 := [...]int{1, 2, 3, 4, 5}
	arr3Sec := arr3[:]
	t.Log(arr3Sec)

}

// 结构体 切片是

// var 变量名 [数组长度]数据类型
//这是一维数组的定义方式，数组长度必须是整数且大于0，未初始化的数组不是nil，也就是说没有空数组（与切片不同）
