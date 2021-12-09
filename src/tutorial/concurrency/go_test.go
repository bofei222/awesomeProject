package concurrency

import (
	"fmt"
	"testing"
	"time"
)

func TestForLoop(t *testing.T) {

	for i := 0; i < 10; i++ {
		go func(val int) {
			fmt.Println(val)
		}(i)
	}
	time.Sleep(1000)

}
