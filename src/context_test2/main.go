package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	AsyncCall()
	fmt.Println(222)
	time.Sleep(20000 * time.Second)
}
func AsyncCall() {
	ctx, cancel := context.WithTimeout(context.Background(), time.Millisecond*800)
	defer cancel()
	go func(ctx context.Context) {
		// 发送HTTP请求
		time.Sleep(time.Millisecond * 900)
		fmt.Println("go end")
	}(ctx)

	select {
	case <-ctx.Done():
		fmt.Println("call successfully!!!")
		return
	case <-time.After(time.Millisecond * 1000):
		fmt.Println("timeout!!!")
		return
	}

}
