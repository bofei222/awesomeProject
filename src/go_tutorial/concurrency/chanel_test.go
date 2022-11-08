package concurrency

import (
	"fmt"
	"testing"
	"time"
)

func sum(s []int, c chan int) {
	sum := 0
	for _, v := range s {
		sum += v
	}
	c <- sum // 将和送入 c

	b := <-c // 消费了 c，
	fmt.Println("b:", b)
	c <- 1 // 填充了c   所以返回时 c 是1
}

func TestChanel(t *testing.T) {
	s := []int{7, 2, 8, -9, 4, 0}

	c := make(chan int)
	go sum(s[:len(s)/2], c)
	go sum(s[len(s)/2:], c)
	x, y := <-c, <-c // 从 c 中接收

	fmt.Println(x, y, x+y)

}

// 主程接收
func TestChanel2(t *testing.T) {
	c := make(chan int)
	go func() {
		c <- 1
	}()
	x := <-c // 从 c 中接收
	fmt.Println(x)
	time.Sleep(time.Second * 10)
}

// 主程填充
func TestChanel3(t *testing.T) {
	c := make(chan int)
	go func() {
		//value := <-c
		//fmt.Println(value)
	}()
	c <- 1 // 向c 中填充
	time.Sleep(time.Second * 10)
}

func TestBufferChanel(t *testing.T) {
	ch := make(chan int, 12)
	ch <- 1
	ch <- 2
	fmt.Println(<-ch)
	fmt.Println(<-ch)
}

func fibonacci(n int, c chan int) {
	x, y := 0, 1
	for i := 0; i < n; i++ {
		c <- x
		x, y = y, x+y
		//fmt.Println(y)
	}
	//close(c)
}

func TestCloseAndRange(t *testing.T) {
	c := make(chan int, 10)
	go fibonacci(cap(c), c)
	for i := range c {
		fmt.Println(i)
	}
}
