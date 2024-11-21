package main

import (
	pb "awesomeProject/proto"
	"context"
	"fmt"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"google.golang.org/grpc"
	"log"
	"math/rand"
	"net/http"
	"time"
)

func main() {
	turbineIDs := make([]string, 1)
	for i := 0; i < 1; i++ {
		turbineIDs[i] = fmt.Sprintf("%04d", i+1)
	}

	// 创建一个 map 用于存储每台风机的客户端连接
	clientMap := make(map[string]pb.WindTurbineServiceClient)

	// 为每台风机创建一个独立的 gRPC 客户端
	for _, turbineID := range turbineIDs {
		conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
		if err != nil {
			log.Fatalf("Failed to connect for turbine %s: %v", turbineID, err)
		}
		defer conn.Close() // 注意：在实际使用时，不要立即关闭连接

		clientMap[turbineID] = pb.NewWindTurbineServiceClient(conn)

		// 启动风机数据发送的 goroutines
		go sendWindTurbineData(clientMap[turbineID], turbineID)
	}

	// 启动 Prometheus 监控
	http.Handle("/metrics", promhttp.Handler())
	http.ListenAndServe(":2112", nil)
	select {} // 保持客户端运行
}

// 发送风机数据
func sendWindTurbineData(client pb.WindTurbineServiceClient, turbineID string) {
	// 定义周期为 1 秒
	const cycleDuration = time.Second / 50

	for {
		// 记录开始时间
		startTime := time.Now()

		// 准备数据
		floatData := make([]float32, 1000)
		boolData := make([]bool, 2000)

		for i := 0; i < 1000; i++ {
			floatData[i] = rand.Float32() * 100.0
		}

		for i := 0; i < 2000; i++ {
			boolData[i] = rand.Intn(2) == 1
		}

		data := &pb.WindTurbineData{
			Timestamp: time.Now().Unix(),
			FloatData: floatData,
			BoolData:  boolData,
			TurbineID: turbineID,
		}

		// 发送数据
		_, err := client.SendData(context.Background(), data)
		if err != nil {
			log.Printf("Failed to send data for turbine %s: %v", turbineID, err)
		} else {
			fmt.Printf("turbine: %s, time: %v\n", turbineID, time.Now())
		}

		// 计算消耗时间
		elapsed := time.Since(startTime)

		// 如果耗时小于周期时间，睡眠补足差值
		if elapsed < cycleDuration {
			time.Sleep(cycleDuration - elapsed)
		}
	}
}
