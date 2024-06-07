package main

import (
	"fmt"
	"time"
)

func main() {

	t := time.NewTimer(1902 * time.Millisecond)
	defer t.Stop()
	go func() {
		select {
		case <-doWork():
			fmt.Println("正常退出")
			return
		case <-t.C:
			fmt.Println("超时")
			return
		}
	}()
	time.Sleep(5 * time.Second)
}
func doWork() chan struct{} {
	var ch = make(chan struct{})

	go func() {
		time.Sleep(1900 * time.Millisecond) //模拟超时
		ch <- struct{}{}
	}()
	return ch
}
