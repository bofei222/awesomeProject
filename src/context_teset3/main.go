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
		// 耗时任务
		time.Sleep(time.Millisecond * 1200)
		fmt.Println("go end")
	}(ctx)

	select {
	case <-ctx.Done():
		fmt.Println("call successfully!!!") // done了 但是，go func 任务并没取消，里面没用ctx做什么操作，比如ctx.Done
		return                              //能返回此方法是因为
	}

}
