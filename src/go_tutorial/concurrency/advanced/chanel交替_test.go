package advanced

import (
	"fmt"
	"testing"
	"time"
)

func TestJiaoTi(t *testing.T) {
	c1 := make(chan int)
	go func() {
		for i := 0; i < 10; i++ {
			if i%2 == 1 { // 1 3 5 7 9
				c1 <- 1
				fmt.Println(i)
			}
		}
	}()
	go func() {
		for i := 0; i < 10; i++ {
			if i%2 == 0 {
				fmt.Println(i)
				<-c1
			}
		}
	}()
	time.Sleep(100 * time.Second)
}
