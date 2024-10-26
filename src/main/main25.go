package main

import (
	"fmt"
	"golang.org/x/net/context"
	"time"
)

const shortDuration = 1050 * time.Millisecond

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), shortDuration)
	defer cancel()

	select {
	case <-time.After(1 * time.Second):
		fmt.Println("脑子进煎鱼了")
	case <-ctx.Done():
		fmt.Println(ctx.Err())
	}

	// a(ctx)

}

// 当传入的 ctx 被取消（通过调用 cancel() 或超时等方式），goroutine 会检测到 ctx.Done() 被关闭，并停止发送数据，退出循环。
// 调用者可以通过读取返回的通道来获取这些连续的整数值
func a(ctx context.Context) <-chan int {
	dst := make(chan int)
	n := 1
	go func() {
		for {
			select {
			case <-ctx.Done():
				return
			case dst <- n:
				n++
			}
		}
	}()
	return dst
}
