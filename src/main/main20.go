package main

import (
	"fmt"
	"sync"
	"time"
)

type Watermark struct {
	timestamp time.Time
}

type Event struct {
	timestamp time.Time
}

type EventStream struct {
	events    []Event
	watermark *Watermark
	mutex     sync.Mutex
	timer     *time.Timer
}

func NewEventStream() *EventStream {
	return &EventStream{
		events:    make([]Event, 0),
		watermark: &Watermark{},
		timer:     time.NewTimer(0),
	}
}

func (s *EventStream) AddEvent(event Event) {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	s.events = append(s.events, event)
	// 重置定时器，确保在新事件到达时立即检查水位线
	s.timer.Reset(0)
}

func (s *EventStream) SetWatermark(watermark Watermark) {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	s.watermark = &watermark
}

func (s *EventStream) ProcessEvents() {
	for {
		select {
		case <-s.timer.C:
			// 定时器触发，处理事件
			s.mutex.Lock()
			for _, event := range s.events {
				if event.timestamp.After(s.watermark.timestamp) {
					// 在水位线之后的事件
					fmt.Printf("Processing event: %v\n", event)
				} else {
					// 在水位线之前的事件，可能需要延迟处理或丢弃
					fmt.Printf("Skipping event: %v (before watermark)\n", event)
				}
			}
			s.events = nil // 清空已处理的事件
			s.mutex.Unlock()

			// 重新设置定时器，等待下一个事件
			s.timer.Reset(time.Second)
		}
	}
}

func main() {
	stream := NewEventStream()

	// 模拟事件流
	go func() {
		for i := 0; i < 10; i++ {
			event := Event{timestamp: time.Now()}
			stream.AddEvent(event)
			time.Sleep(500 * time.Millisecond) // 模拟事件之间的时间间隔
		}
	}()

	// 模拟水位线的更新
	go func() {
		for {
			time.Sleep(3 * time.Second) // 模拟水位线更新的时间间隔
			watermark := Watermark{timestamp: time.Now()}
			stream.SetWatermark(watermark)
			fmt.Printf("Watermark updated: %v\n", watermark)
		}
	}()

	// 处理事件
	stream.ProcessEvents()

	// 防止主goroutine退出
	select {}
}
