package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), time.Millisecond*800)
	defer cancel()

	go AsyncCall(ctx)

	time.Sleep(20000 * time.Second)
}

func AsyncCall(ctx context.Context) {

	fmt.Println(222)
	select {
	default:
		fmt.Println(444)
	case <-ctx.Done():
		fmt.Println("call successfully!!!") // done了 但是，go func 任务并没取消，里面没用ctx做什么操作，比如ctx.Done
		return
	}
}
