package asleep

import (
	"fmt"
	"testing"
)

func Test(t *testing.T) {
	ch1 := make(chan string)
	ch1 <- "hello world"
	fmt.Println(<-ch1)
}

func Test2(t *testing.T) {
	ch1 := make(chan string)
	go func() {
		ch1 <- "hello world"
	}()
	fmt.Println(<-ch1)
}

// 与Test相比对了缓冲
func TestBufferChanel(t *testing.T) {
	ch1 := make(chan string, 1)
	ch1 <- "hello world"
	fmt.Println(<-ch1)
}

func TestMultiValueChanel(t *testing.T) {
	ch1 := make(chan string)
	go func() {
		fmt.Println(<-ch1)
	}()
	ch1 <- "hello world"
	ch1 <- "hello China"
}

// 不会出现deadlock
func TestMultiValueChanel2(t *testing.T) {
	ch1 := make(chan string)
	go func() {
		ch1 <- "hello world"
		ch1 <- "hello China"
	}()
	fmt.Println(<-ch1)

}
