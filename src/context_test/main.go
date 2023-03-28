package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	// 创建一个子节点的context,3秒后自动超时
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)

	go watch(ctx, "监控1")
	go watch(ctx, "监控2")

	fmt.Println("现在开始等待8秒,time=", time.Now())
	time.Sleep(8 * time.Second)

	fmt.Println("等待8秒结束,准备调用cancel()函数，发现两个子协程已经结束了，time=", time.Now())
	cancel()
}

// 单独的监控协程
func watch(ctx context.Context, name string) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println(name, "收到信号，监控退出,time=", time.Now())
			return
		default:
			fmt.Println(name, "goroutine监控中,time=", time.Now())
			time.Sleep(4 * time.Second) // 如果被堵塞在这里是退出不了的，
			fmt.Println(name, "goroutine监控结束,time=", time.Now())
		}
	}
}
