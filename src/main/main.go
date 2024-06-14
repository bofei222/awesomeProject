package main

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	"github.com/goburrow/modbus"
	"log"
	"os"
	"time"
)

const (
	//rtuDevice = "/dev/ttyUSB0" TODO 蓝色
	//rtuDevice = "/dev/ttyUSB0"  TODO 黑色
	//rtuDevice = "COM4" // TODO 蓝色
	rtuDevice = "COM5" // TODO 黑色
)

// 全局变量 context，用于 Redis 操作
var ctx2 = context.Background()

func main() {

	handler := modbus.NewRTUClientHandler(rtuDevice)
	handler.BaudRate = 9600
	handler.DataBits = 8
	handler.Parity = "N"
	handler.StopBits = 1
	handler.SlaveId = 1
	handler.Logger = log.New(os.Stdout, "rtu: ", log.LstdFlags)
	err := handler.Connect()
	if err != nil {
		log.Fatal(err)
	}
	defer handler.Close()

	client := modbus.NewClient(handler)

	// 创建 Redis 客户端
	rdb := redis.NewClient(&redis.Options{
		Addr:     "192.168.1.244:6379", // Redis 服务器地址和端口 //TODO
		Password: "zTGBBRP3dx2VG8su",   // Redis 密码
		DB:       2,                    // Redis 数据库编号
	})

	// 定时器，每2秒执行一次
	ticker := time.NewTicker(2 * time.Second)
	defer ticker.Stop()

	var lastExists bool

	for range ticker.C {
		// 查询 Redis 中的 hash 键 eventCenter:prod:light
		exists, err := rdb.Exists(ctx2, "eventCenter:prod:light").Result()
		if err != nil {
			fmt.Println("查询 Redis 错误：", err)
			continue
		}

		currentExists := exists > 0

		// 检查 hash 键是否从存在变为不存在
		if lastExists && !currentExists {
			fmt.Println("键从存在变为不存在,发送开启指令")
			results, err := client.WriteSingleRegister(194, 6)
			if err != nil || results == nil {
				log.Fatal(err, results)
			}
		} else if !lastExists && currentExists {
			fmt.Println("键从不存在变为存在,发送关闭指令")
			results, err := client.WriteSingleRegister(194, 4)
			if err != nil || results == nil {
				log.Fatal(err, results)
			}
		}

		lastExists = currentExists

		// //results, err := client.WriteSingleRegister(17, 1) 全部循环和关闭
		//	//results, err = client.WriteSingleRegister(22, 1)
	}
}
