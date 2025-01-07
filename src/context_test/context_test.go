package main

import (
	"context"
	"fmt"
	"sync"
	"testing"
	"time"
)

func TestTimeout(t *testing.T) {
	wg := sync.WaitGroup{}
	wg.Add(1)

	parentContext := context.Background()
	cancelCtx, cancelFunc := context.WithTimeout(parentContext, time.Second*3)
	go func(ctx context.Context) {
		t.Log("goroutine 1")
		select {
		case <-cancelCtx.Done():
			fmt.Println("bbb")
			wg.Done()
			return
		case <-time.After(time.Hour * 100):
			fmt.Println("aaa")
			wg.Done()
		}

	}(cancelCtx)
	wg.Wait()
	cancelFunc()
	t.Log("goroutine cgo")
}
