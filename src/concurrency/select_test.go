package concurrency

import (
	"fmt"
	"testing"
)

func fibonacci2(c, quit chan int) {
	x, y := 0, 1
	for {
		fmt.Println("+")
		select {
		case c <- x:
			x, y = y, x+y
		case <-quit:
			fmt.Println("quit")
			return
		}
	}
}

// 没有给长度（一个不存 一个不缓冲） ， 为什么消费10次没有报错， - + 是交次进行的
func TestSelect(t *testing.T) {
	c := make(chan int)
	fmt.Println(cap(c))
	quit := make(chan int)
	go func() {
		for i := 0; i < 10; i++ {
			fmt.Println("-")
			fmt.Println(<-c)
		}
		quit <- 0
	}()
	fibonacci2(c, quit)
}
