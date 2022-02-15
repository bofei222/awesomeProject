package customer_type

import (
	"fmt"
	"time"
)

// 自定义类型
//方法声明太长
// 方法定义 转化为 一个类型  ，别名

type IntConv func(op int) int

func timeSpent(inner IntConv) IntConv {
	return func(op int) int {
		start := time.Now().UTC()
		ret := inner(op)
		fmt.Println("time spend:", time.Since(start).Seconds())
		return ret
	}
}
