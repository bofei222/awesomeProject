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
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Failed to connect: %v", err)
	}
	defer conn.Close()

	client := pb.NewWindTurbineServiceClient(conn)
	turbineIDs := make([]string, 300)
	for i := 0; i < 300; i++ {
		turbineIDs[i] = fmt.Sprintf("%04d", i+1)
	}

	// 启动风机数据发送的 goroutines
	for _, turbineID := range turbineIDs {
		go sendWindTurbineData(client, turbineID)
	}

	/*// 定时查询风机 0003 的平均值
	go func() {
		for {
			time.Sleep(30 * time.Second)
			getWindTurbineAverage(client, "0003")
		}
	}()

	// 定时查询全场风机的平均值
	go func() {
		for {
			time.Sleep(3 * time.Minute)
			getAllWindTurbinesAverage(client)
		}
	}()*/
	http.Handle("/metrics", promhttp.Handler())
	http.ListenAndServe(":2112", nil)
	select {} // 保持客户端运行
}

// 发送风机数据
func sendWindTurbineData(client pb.WindTurbineServiceClient, turbineID string) {
	for {

		time.Sleep(time.Millisecond * 1000)
		floatData := make([]float32, 1000)
		boolData := make([]bool, 1000)

		for i := 0; i < 1000; i++ {
			floatData[i] = rand.Float32() * 100.0
			boolData[i] = rand.Intn(2) == 1
		}

		data := &pb.WindTurbineData{
			Timestamp: time.Now().Unix(),
			FloatData: floatData,
			BoolData:  boolData,
			TurbineID: turbineID,
		}

		_, err := client.SendData(context.Background(), data)
		if err != nil {
			log.Printf("Failed to send data for turbine %s: %v", turbineID, err)
		}
	}

}

// 查询单个风机的平均值
func getWindTurbineAverage(client pb.WindTurbineServiceClient, turbineID string) {
	req := &pb.WindTurbineAverageRequest{TurbineID: turbineID}
	resp, err := client.GetWindTurbineAverage(context.Background(), req)
	if err != nil {
		log.Printf("Failed to get average for turbine %s: %v", turbineID, err)
		return
	}
	fmt.Printf("Average for turbine %s: %f\n", turbineID, resp.Average)
}

// 查询全场风机的平均值
func getAllWindTurbinesAverage(client pb.WindTurbineServiceClient) {
	req := &pb.AllWindTurbinesAverageRequest{}
	resp, err := client.GetAllWindTurbinesAverage(context.Background(), req)
	if err != nil {
		log.Printf("Failed to get all turbines average: %v", err)
		return
	}
	fmt.Printf("All turbines average: %f\n", resp.Average)
}
