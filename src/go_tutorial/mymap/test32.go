package main

import "fmt"

func main() {
	m := make(map[string]int)

	m["Answer"] = 42
	fmt.Println("The value:", m["Answer"])

	m["Answer"] = 48
	fmt.Println("The value:", m["Answer"])

	delete(m, "Answer")
	fmt.Println("The value:", m["Answer"])

	v, ok := m["Answer"]
	fmt.Println("The value:", v, "Present?", ok)

	x := m["Answer"]
	fmt.Println(x)
	y := m["Answer"]
	fmt.Println(y)
	// elem, ok = m[key] 通过双赋值检测某个键是否存在：若 key 不在映射中，那么 elem 是该映射元素类型的零值。
	a, b := m["Answer"]
	fmt.Println(a)
	fmt.Println(b)

}
