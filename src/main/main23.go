package main

import (
	"context"
	"fmt"
	"time"

	"github.com/go-redis/redis/v8"
)

// 全局变量 context，用于 Redis 操作
var ctx = context.Background()

func main() {
	// 创建 Redis 客户端
	rdb := redis.NewClient(&redis.Options{
		Addr:     "192.168.1.244:6379", // Redis 服务器地址和端口
		Password: "zTGBBRP3dx2VG8su",   // Redis 密码，如果没有则留空
		DB:       2,                    // Redis 数据库编号
	})

	// 定时器，每2秒执行一次
	ticker := time.NewTicker(2 * time.Second)
	defer ticker.Stop()

	for range ticker.C {
		// 查询 Redis 中的 hash 键 eventCenter:prod:light
		exists, err := rdb.Exists(ctx, "eventCenter:prod:light").Result()
		if err != nil {
			fmt.Println("查询 Redis 错误：", err)
			continue
		}

		// 检查 hash 键是否存在
		if exists > 0 {
			fmt.Println("存在")
		} else {
			fmt.Println("不存在")
		}
	}
}
