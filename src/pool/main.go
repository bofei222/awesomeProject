package main

import (
	"fmt"
	"github.com/panjf2000/ants/v2"
	"sync"
	"time"
)

func wrapper(i int, wg *sync.WaitGroup) func() {
	return func() {
		fmt.Printf("hello from task:%d\n", i)
		time.Sleep(3 * time.Second)
		wg.Done()
	}
}

func main() {
	p, _ := ants.NewPool(4, ants.WithMaxBlockingTasks(100))
	defer p.Release()

	var wg sync.WaitGroup
	wg.Add(8)
	for i := 1; i <= 8; i++ {
		go func(i int) {
			err := p.Submit(wrapper(i, &wg))
			if err != nil {
				fmt.Printf("task:%d err:%v\n", i, err)
				wg.Done()
			}
		}(i)
	}
	fmt.Println("All tasks are submit!")
	wg.Wait()
	fmt.Println("All tasks are done!")
}
