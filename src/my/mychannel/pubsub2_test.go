package mychannel

import (
	"fmt"
	"strings"
	"sync"
	"testing"
	"time"
)

type (
	subscriber chan interface{}         //订阅者 信道
	filter     func(v interface{}) bool //过滤器
)

// 信息发布者
type Publisher struct {
	m           sync.RWMutex          //读写锁
	buffer      int                   //一个订阅者的缓存信息数量
	timeout     time.Duration         //信息的延迟时间
	subscribers map[subscriber]filter //订阅者
}

// 获取一个信息发布者
func NewPublisher(duration time.Duration, int2 int) *Publisher {
	return &Publisher{
		buffer:      int2,
		timeout:     duration,
		subscribers: make(map[subscriber]filter), //映射初始化
	}
}

// 信息发布者关闭
func (p *Publisher) Close() {
	p.m.Lock()
	defer p.m.Unlock()

	for sub, _ := range p.subscribers {
		delete(p.subscribers, sub)
		close(sub)
	}
}

// 给所有的人发布信息
func (p *Publisher) PublishAll(info interface{}) {
	p.m.RLock()
	defer p.m.RUnlock()

	var wg sync.WaitGroup //此时，channel只是一个接收数据的东西
	for sub, filt := range p.subscribers {
		wg.Add(1)
		go p.PublishOne(sub, filt, info, &wg) //给一个人发放信息
	}

	wg.Wait()
	return
}

// 给一个人发布信息
func (p *Publisher) PublishOne(sub subscriber, filt filter, info interface{}, wg *sync.WaitGroup) {
	defer wg.Done()

	if filt != nil && !filt(info) {
		return
	}

	select {
	case sub <- info:
	case <-time.After(p.timeout):
	}
}

// 获取信息订阅者
func (p *Publisher) GetSubscriber(filter2 filter) chan interface{} {
	p.m.Lock()
	defer p.m.Unlock()

	ch := make(chan interface{}, p.buffer)
	p.subscribers[ch] = filter2

	return ch
}

// 删除一个信息订阅者
func (p *Publisher) DelSubscriber(sub subscriber) {
	p.m.Lock()
	defer p.m.Unlock()

	delete(p.subscribers, sub)
	close(sub)
}

func TestA(t *testing.T) {
	//获取信息发布者
	p := NewPublisher(100*time.Microsecond, 32)
	defer p.Close()

	//获取订阅者1
	sub1 := p.GetSubscriber(nil)
	//获取订阅者2
	sub2 := p.GetSubscriber(func(v interface{}) bool {
		if str, ok := v.(string); ok {
			return strings.Contains(str, "gongyao")
		}
		return false
	})

	//发布信息
	p.PublishAll("Hello gongyao")
	p.PublishAll("Hello wanghui")

	go func() { //读取 订阅者1受到的消息
		for sub1_msg := range sub1 {
			fmt.Println("Sub1 : ", sub1_msg)
		}
	}()
	go func() { //读取 订阅者2受到的消息
		for sub2_msg := range sub2 {
			fmt.Println("Sub2 : ", sub2_msg)
		}
	}()

	time.Sleep(1 * time.Second)
}
