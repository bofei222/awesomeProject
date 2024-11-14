package main

import (
	"fmt"
	"github.com/panjf2000/ants/v2"
	"time"
)

func main() {
	// 创建一个最大并发数为 3 的 Goroutine 池
	pool, err := ants.NewPool(3)
	if err != nil {
		fmt.Println("Failed to create pool:", err)
		return
	}
	defer pool.Release() // 确保池在程序结束时释放资源

	// 提交 10 个任务到池中
	for i := 0; i < 10; i++ {
		// 每个任务会模拟一个耗时 1 秒的操作
		taskID := i
		pool.Submit(func() {
			fmt.Printf("Task %d is starting\n", taskID)
			time.Sleep(1 * time.Second) // 模拟任务执行
			fmt.Printf("Task %d is done\n", taskID)
		})
	}

	// 等待一段时间让所有任务完成
	time.Sleep(30 * time.Second) // 等待任务完成，确保所有任务都执行完
	fmt.Println("All tasks are done!")
}
