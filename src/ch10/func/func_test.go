package _func

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

func returnMultiValues() (int, int) {
	return rand.Intn(10), rand.Intn(20)
}

//在原来的函数上 包了一下
// function programming
func timeSpent(inner func(op int) int) func(op int) int {
	return func(op int) int {
		start := time.Now().UTC()
		ret := inner(op)
		fmt.Println("time spend:", time.Since(start).Seconds())
		return ret
	}
}
func slowFun(op int) int {
	time.Sleep(time.Second * 2)
	return op
}
func TestFun(t *testing.T) {
	a, _ := returnMultiValues()
	t.Log(a)

	tsSF := timeSpent(slowFun) //没有调用slowFun,只是操作值一样操作
	tsSF(a)
}

// 计算机程序的构造和解析  MIT

//科扁参数
func Sum(ops ...int) {

}
func TestVarParam(t *testing.T) {

}

// 延迟函数 defer 清理资源 释放锁 panic 程序中断，后面代码不能执行  defer匿名函数 非匿名函数都可
