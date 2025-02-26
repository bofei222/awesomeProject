package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int)

	// 一个 goroutine 不断写数据到 channel
	go func() {
		for i := 0; i < 10; i++ {
			ch <- i
			fmt.Println("Sent:", i)
			time.Sleep(100 * time.Millisecond)
		}
		close(ch)
	}()

	// 主 goroutine 读取数据
	for val := range ch {
		fmt.Println("Received:", val)
	}

	// 漏掉的部分：没有处理 channel 阻塞的情况。
	// 当所有 goroutine 都结束时，channel 可能因为没有人读取而导致写入阻塞
}
