package csp

import (
	"fmt"
	"testing"
	"time"
)

func service() string {
	time.Sleep(time.Millisecond * 50)
	return "Done"
}

func otherTask() {
	fmt.Println("other start")
	time.Sleep(time.Millisecond * 100)
	fmt.Println("other end")
}

func TestService(t *testing.T) {
	fmt.Println(service())
	otherTask()
}

func AsyncService() chan string {
	retCh := make(chan string)
	go func() {
		ret := service()
		fmt.Println("returned result")
		retCh <- ret // 有点像 Future
		fmt.Println("service exited")
	}()
	return retCh
}

func TestAsyncService(t *testing.T) {
	retCh := AsyncService()
	otherTask()
	fmt.Println(<-retCh) // 类似 future.get 堵塞直至子线程完成
}

func TestAsyncService2(t *testing.T) {
	retCh := AsyncService()
	fmt.Println(<-retCh) // 在这里 堵塞 耗时会是 ：0.15s
	otherTask()
}

func TestAsyncService3(t *testing.T) {
	go service() // 不能得到协程中的 返回值
	otherTask()
}
