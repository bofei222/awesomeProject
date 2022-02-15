package geek

import (
	"fmt"
	"testing"
	"time"
)

func TestChan1(t *testing.T) {
	// Creating a channel
	// Using var keyword
	var mychannel chan int
	fmt.Println("Value of the channel: ", mychannel)
	fmt.Printf("Type of the channel: %T ", mychannel)

	// Creating a channel using make() function
	mychannel1 := make(chan int)
	fmt.Println("\nValue of the channel1: ", mychannel1)
	fmt.Printf("Type of the channel1: %T ", mychannel1)

}

func myfunc(ch chan int) {
	fmt.Println(time.Now())
	a := <-ch // 此goroutine 会堵塞 直到主线程 填充值到通道
	//b := <-ch
	//fmt.Println("b:",b)
	fmt.Println(time.Now())
	fmt.Println(234 + a)
}

// 通道 自带 列表属性
func TestChan2(t *testing.T) {
	fmt.Println("start Main method")
	// Creating a channel
	ch := make(chan int)
	go myfunc(ch)
	time.Sleep(time.Second * 2)
	ch <- 23
	fmt.Println("End Main method")

	select {} // 阻塞

}

// Function
func myfun(mychnl chan string) {

	for v := 0; v < 3; v++ {
		mychnl <- "GeeksforGeeks"
	}
	close(mychnl) //fatal error: all goroutines are asleep - deadlock! 注释掉会报错
}

//  为什么不关闭  继续读会报错 而不是堵塞直接
func TestChanClose(t *testing.T) {
	// Creating a channel
	c := make(chan string)

	// calling Goroutine
	go myfun(c)

	// When the value of ok is
	// set to true means the
	// channel is open and it
	// can send or receive data
	// When the value of ok is set to
	// false means the channel is closed
	for {
		res, ok := <-c
		if ok == false {
			fmt.Println("Channel Close ", ok)
			break
		}
		fmt.Println("Channel Open ", res, ok)
	}
}

func TestRangeChnl(t *testing.T) {
	// Creating a channel
	// Using make() function
	mychnl := make(chan string, 4)

	// Anonymous goroutine
	go func() {
		mychnl <- "GFG"
		mychnl <- "gfg"
		mychnl <- "Geeks"
		mychnl <- "GeeksforGeeks"
		close(mychnl)
	}()

	// Using for loop
	fmt.Println(len(mychnl))
	for res := range mychnl {
		fmt.Println(res, len(mychnl))
	}
}

func TestChnlLen(t *testing.T) {
	// Creating a channel
	// Using make() function
	mychnl := make(chan string, 4)
	mychnl <- "GFG"
	mychnl <- "gfg"
	mychnl <- "Geeks"
	mychnl <- "GeeksforGeeks"

	// Finding the length of the channel
	// Using len() function
	fmt.Println(<-mychnl)
	fmt.Println("Length of the channel is: ", len(mychnl))

}

func TestChannelCap(t *testing.T) {
	// Creating a channel
	// Using make() function
	mychnl := make(chan string, 5)
	mychnl <- "GFG"
	mychnl <- "gfg"
	mychnl <- "Geeks"
	mychnl <- "GeeksforGeeks"

	// Finding the capacity of the channel
	// Using cap() function
	fmt.Println("Capacity of the channel is: ", cap(mychnl))
}
