package ch7

import "testing"

// map的 value 可以是 函数
func TestMapWithFunValue(t *testing.T) {
	m := map[int]func(op int) int{}
	m[1] = func(op int) int { return op }
	m[2] = func(op int) int { return op * op }
	m[3] = func(op int) int { return op * op * op }
	//t.Log(m[1]) Log arg m[1] is a func value, not called
	t.Log(m[1](2), m[2](2), m[3](2))

}

// 实现set
func TestMapForSet(t *testing.T) {
	mySet := map[int]bool{}
	mySet[1] = true

	// 利用bool 类型默认是false的特性1)
	n := 3
	if mySet[n] {
		t.Logf("%d exist", n)
	} else {
		t.Logf("%d not exist", n)
	}
	mySet[3] = true
	t.Log(len(mySet))
	delete(mySet, 1)
	n = 1
	if mySet[n] {
		t.Logf("%d exist", n)
	} else {
		t.Logf("%d not exist", n)
	}

}
