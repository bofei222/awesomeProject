package share_mem

import (
	"fmt"
	"sync"
	"sync/atomic"
	"testing"
	"time"
)

func TestCounter(t *testing.T) {
	counter := 0
	for i := 0; i < 5000; i++ {
		go func() {
			counter = counter + 1
		}()
	}
	//counter++  不是线程安全的 就算睡1s 结果也不是5000
	time.Sleep(time.Second * 1)
	fmt.Println(counter)
}

// Mutex
func TestCounterThreadSafe(t *testing.T) {
	counter := 0
	var lock sync.Mutex
	for i := 0; i < 5000; i++ {
		go func() {

			defer func() {
				lock.Unlock()
			}()
			lock.Lock()
			counter++
		}()
	}
	time.Sleep(time.Second * 1) // lock不会等待线程完成，只是保证计数线程安全
	fmt.Println(counter)
}

// waitGroup 独写锁分开  ,本身没有 lock功能
func TestCounterWaitGroup(t *testing.T) {
	counter := 0
	var wg sync.WaitGroup
	var lock sync.RWMutex // 为什么声明后不用初始化
	for i := 0; i < 5000; i++ {
		wg.Add(1) // 像 java 计数器 DownCount  ,join
		go func() {
			defer func() {
				lock.Unlock()
			}()
			lock.Lock()
			counter++
			wg.Done()
		}()
	}
	wg.Wait()
	//time.Sleep(time.Second*1)
	t.Log(counter)

}

func TestAuto(t *testing.T) {
	var ops uint64
	var wg sync.WaitGroup
	for i := 0; i < 50; i++ {
		wg.Add(1)
		go func() {
			for c := 0; c < 1000; c++ {
				atomic.AddUint64(&ops, 1)
			}
			defer func() {
				wg.Done()
			}()
		}()
	}
	wg.Wait()
	fmt.Println("ops:", ops)
}
