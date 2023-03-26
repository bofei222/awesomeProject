package mychannel

import (
	"fmt"
	"testing"
)

func TestRangeChannel(t *testing.T) {
	// 初始化channel
	ch := make(chan int, 5)
	go func() {
		for i := 0; i < 5; i++ {
			ch <- i
		}
	}()
	for i := range ch {
		fmt.Println(i)
	}
}
