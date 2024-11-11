package main

import (
	"context"
	"fmt"
	"log"
	_ "math/rand"
	"net"
	"net/http"
	"sync"
	"time"

	pb "awesomeProject/proto" // 根据生成的 proto 文件路径修改
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"google.golang.org/grpc"
)

// Prometheus 指标
var (
	requestCount = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "wind_turbine_requests_total",
			Help: "Total number of requests received by the wind turbine service",
		},
		[]string{"method"},
	)
	requestDuration = prometheus.NewHistogramVec(
		prometheus.HistogramOpts{
			Name:    "wind_turbine_request_duration_seconds",
			Help:    "Histogram of request durations for the wind turbine service",
			Buckets: prometheus.DefBuckets,
		},
		[]string{"method"},
	)
	memoryUsage = prometheus.NewGauge(
		prometheus.GaugeOpts{
			Name: "wind_turbine_memory_usage_bytes",
			Help: "Current memory usage of the wind turbine service",
		},
	)
)

// 记录所有风机数据的内存存储
var windTurbineData = make(map[string][]*pb.WindTurbineData)
var mu sync.Mutex

// 服务端结构体
type WindTurbineServer struct {
	pb.UnimplementedWindTurbineServiceServer
}

// 初始化 Prometheus 指标
func init() {
	prometheus.MustRegister(requestCount)
	prometheus.MustRegister(requestDuration)
	prometheus.MustRegister(memoryUsage)
}

// 处理数据写入
func (s *WindTurbineServer) SendData(ctx context.Context, data *pb.WindTurbineData) (*pb.WriteResponse, error) {
	// 记录请求开始时间
	start := time.Now()

	mu.Lock()
	defer mu.Unlock()

	windTurbineData[data.TurbineID] = append(windTurbineData[data.TurbineID], data)

	// 更新指标
	requestCount.WithLabelValues("SendData").Inc()
	requestDuration.WithLabelValues("SendData").Observe(time.Since(start).Seconds())

	return &pb.WriteResponse{Message: "Data received"}, nil
}

// 查询单个风机的平均值
func (s *WindTurbineServer) GetWindTurbineAverage(ctx context.Context, req *pb.WindTurbineAverageRequest) (*pb.WindTurbineAverageResponse, error) {
	// 记录请求开始时间
	start := time.Now()

	mu.Lock()
	defer mu.Unlock()

	data := windTurbineData[req.TurbineID]
	if len(data) == 0 {
		return &pb.WindTurbineAverageResponse{Average: 0}, nil
	}

	var sum float32
	var count int
	for _, entry := range data {
		for _, value := range entry.FloatData {
			sum += value
			count++
		}
	}

	// 更新指标
	requestCount.WithLabelValues("GetWindTurbineAverage").Inc()
	requestDuration.WithLabelValues("GetWindTurbineAverage").Observe(time.Since(start).Seconds())

	return &pb.WindTurbineAverageResponse{Average: sum / float32(count)}, nil
}

// 查询所有风机的全场平均值
func (s *WindTurbineServer) GetAllWindTurbinesAverage(ctx context.Context, req *pb.AllWindTurbinesAverageRequest) (*pb.WindTurbinesAverageResponse, error) {
	// 记录请求开始时间
	start := time.Now()

	mu.Lock()
	defer mu.Unlock()

	var totalSum float32
	var totalCount int
	for _, dataList := range windTurbineData {
		for _, entry := range dataList {
			for _, value := range entry.FloatData {
				totalSum += value
				totalCount++
			}
		}
	}

	// 更新指标
	requestCount.WithLabelValues("GetAllWindTurbinesAverage").Inc()
	requestDuration.WithLabelValues("GetAllWindTurbinesAverage").Observe(time.Since(start).Seconds())

	return &pb.WindTurbinesAverageResponse{Average: totalSum / float32(totalCount)}, nil
}

// 用于监控的 HTTP 端点
func startMetricsServer() {
	http.Handle("/metrics", promhttp.Handler())
	fmt.Println("Starting Prometheus metrics server on :9090")
	log.Fatal(http.ListenAndServe(":9090", nil))
}

func main() {
	// 启动 Prometheus 监控 HTTP 服务器
	go startMetricsServer()

	listener, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("Failed to listen on port 50051: %v", err)
	}

	grpcServer := grpc.NewServer()
	pb.RegisterWindTurbineServiceServer(grpcServer, &WindTurbineServer{})
	fmt.Println("Server is running on port :50051")

	// 启动 gRPC 服务器
	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
