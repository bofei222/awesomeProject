package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int)

	// 启动一个 goroutine 写数据
	go func() {
		for i := 0; i < 10; i++ {
			ch <- i
			fmt.Println("Sent:", i)
			time.Sleep(100 * time.Millisecond)
		}
		close(ch)
	}()

	// 主 goroutine 读取数据，确保读取完成后退出
	for {
		select {
		case val, ok := <-ch:
			if !ok {
				// channel 已关闭且所有数据已读完
				fmt.Println("Channel closed!")
				return
			}
			fmt.Println("Received:", val)
		}
	}
}
