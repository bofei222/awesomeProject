package t2

import (
	"fmt"
	"sync"
	"testing"
)

var wg = sync.WaitGroup{}

func TestGoroutine1(t *testing.T) {
	userCount := 10
	for i := 0; i < userCount; i++ {
		wg.Add(1)
		go Read(i)
	}

	wg.Wait()
}

func Read(i int) {
	defer wg.Done()
	fmt.Printf("go func: %d\n", i)
}
