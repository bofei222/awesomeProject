package _interface

import (
	"fmt"
	"testing"
)

// java  显示实现  可 包循环依赖

type Programmer interface {
	WriteHelloWorld() string
}

// 完全看不出 实现
// 看起来是鸭子 就是鸭子
type GoProgrammer struct {
}

// 签名完全一直
// 可以先 写方法 ，之后 在抽象出接口
func (g *GoProgrammer) WriteHelloWorld() string {
	fmt.Println("Hello World Go!")
	return "Hello World Go"
}

func TestHelloWorld(t *testing.T) {
	var programmer Programmer
	programmer = new(GoProgrammer)
	t.Log(programmer.WriteHelloWorld())
}
