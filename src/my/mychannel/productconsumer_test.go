package mychannel

import (
	"fmt"
	"testing"
	"time"
)

// 生产者
func Producer(ch chan int) {
	for i := 1; ; i++ {
		ch <- i
	}
}

// 生产者
func Producer2() chan int {
	ch := make(chan int)

	go func() {
		for i := 0; ; i++ {
			ch <- i
		}
	}()

	return ch
}

// 消费者
func Consumer(ch chan int) {
	for value := range ch {
		fmt.Println(value)
	}
}
func TestProductConsumer(t *testing.T) {
	ch := make(chan int, 64)
	go Producer(ch)
	go Consumer(ch)
	time.Sleep(2 * time.Second)
}

func TestChannel(t *testing.T) {
	ch := Producer2()
	go Consumer(ch)
	time.Sleep(2 * time.Second)
}
