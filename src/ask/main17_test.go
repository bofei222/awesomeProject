package ask

import (
	"fmt"
	"testing"
	"time"
)

func TestMain17(t *testing.T) {
	//初始化定时器
	t2 := time.NewTimer(2 * time.Second)
	//当前时间
	now := time.Now()
	fmt.Printf("Now time : %v.\n", now)

	// C是一个chan time.Time类型的缓冲通道，一旦触及到期时间，定时器就会向自己的C字段发送一个time.Time类型的元素值
	expire := <-t2.C
	fmt.Printf("Expiration time: %v.\n", expire)
}

func Test18(t *testing.T) {
	ch1 := make(chan int, 1)
	select {
	case e1 := <-ch1:
		//如果ch1通道成功读取数据，则执行该case处理语句
		fmt.Printf("1th case is selected. e1=%v", e1)
	case <-time.After(2 * time.Second):
		fmt.Println("Timed out")
	}
}
func Test19(t *testing.T) {
	f := func() {
		fmt.Printf("Expiration time : %v.\n", time.Now())
	}
	time.AfterFunc(1*time.Second, f)
	time.Sleep(2 * time.Second)
}
