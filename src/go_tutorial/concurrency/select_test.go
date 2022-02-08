package concurrency

import (
	"fmt"
	"testing"
	"time"
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

func TestTimeTick(t *testing.T) {
	fmt.Println(time.Now().Second())
	fmt.Println("sanyWind")
	select {
	case <-time.Tick(2 * time.Second):
		fmt.Println("2 second over:", time.Now().Second())
	case <-time.After(7 * time.Second):
		fmt.Println("5 second over, timeover", time.Now().Second())
		return
	}
}

// 当 select 中的其它分支都没有准备好时，default 分支就会执行。
// 为了在尝试发送或者接收时不发生阻塞，可使用 default 分支
func TestDefaultSelect(t *testing.T) {
	tick := time.Tick(100 * time.Millisecond)
	boom := time.After(500 * time.Millisecond)
	for {
		select {
		case <-tick:
			fmt.Println("tick.")
		case <-boom:
			fmt.Println("BOOM!")
			return
		default:
			fmt.Println("    .")
			time.Sleep(50 * time.Millisecond)
		}
	}
}
