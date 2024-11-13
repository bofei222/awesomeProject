package main

import (
	"context"
	"fmt"
	"log"
	_ "math/rand"
	"net"
	"net/http"
	"os"
	"sync"
	"time"

	pb "awesomeProject/proto" // 根据生成的 proto 文件路径修改
	"github.com/apache/arrow-go/v18/arrow"
	"github.com/apache/arrow-go/v18/arrow/array"
	"github.com/apache/arrow-go/v18/arrow/memory"
	"github.com/apache/arrow-go/v18/parquet/pqarrow"
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

// 初始化 Prometheus 指标
func init() {
	prometheus.MustRegister(requestCount)
	prometheus.MustRegister(requestDuration)
	prometheus.MustRegister(memoryUsage)
}

// 服务端结构体
type WindTurbineServer struct {
	pb.UnimplementedWindTurbineServiceServer
}

// WindTurbineData 定义了单条数据的结构
type WindTurbineData struct {
	Timestamp int64     // 时间戳，粒度为秒
	FloatData []float32 // 1000 个 float32 点位
	BoolData  []bool    // 1000 个 bool 点位
}

// createArrowTable 创建 Arrow Table，将每台风机的数据转换为 Arrow 格式
func createArrowTable(turbineData []*pb.WindTurbineData) (*arrow.Table, error) {
	// 构建 Schema
	fields := []arrow.Field{
		{Name: "Timestamp", Type: arrow.PrimitiveTypes.Int64},
	}

	// 创建 1000 个 float32 列
	for i := 0; i < 1000; i++ {
		fields = append(fields, arrow.Field{Name: fmt.Sprintf("MC%04d", i+1), Type: arrow.PrimitiveTypes.Float32})
	}

	// 创建 1000 个 bool 列
	for i := 0; i < 1000; i++ {
		fields = append(fields, arrow.Field{Name: fmt.Sprintf("MA%04d", i+1), Type: new(arrow.BooleanType)})
	}

	// 创建 Schema
	schema := arrow.NewSchema(fields, nil)

	// 创建 Builder 和 Chunked 数组
	timestampBuilder := array.NewInt64Builder(memory.DefaultAllocator)
	floatBuilders := make([]*array.Float32Builder, 1000)
	boolBuilders := make([]*array.BooleanBuilder, 1000)

	// 初始化 float32 和 bool 构建器
	for i := 0; i < 1000; i++ {
		floatBuilders[i] = array.NewFloat32Builder(memory.DefaultAllocator)
		boolBuilders[i] = array.NewBooleanBuilder(memory.DefaultAllocator)
	}

	// 填充数据
	for _, data := range turbineData {
		timestampBuilder.Append(data.Timestamp)

		for i := 0; i < 1000; i++ {
			floatBuilders[i].Append(data.FloatData[i])
			boolBuilders[i].Append(data.BoolData[i])
		}
	}

	// 创建 Chunked 数组
	timestampChunk := arrow.NewChunked(arrow.PrimitiveTypes.Int64, []arrow.Array{timestampBuilder.NewArray()})
	floatChunks := make([]arrow.Chunked, 1000)
	boolChunks := make([]arrow.Chunked, 1000)

	for i := 0; i < 1000; i++ {
		floatChunks[i] = *arrow.NewChunked(arrow.PrimitiveTypes.Float32, []arrow.Array{floatBuilders[i].NewArray()})
		boolChunks[i] = *arrow.NewChunked(new(arrow.BooleanType), []arrow.Array{boolBuilders[i].NewArray()})
	}

	// 创建 Arrow Column
	columns := []arrow.Column{
		*arrow.NewColumn(schema.Field(0), timestampChunk),
	}

	for i := 0; i < 1000; i++ {
		columns = append(columns, *arrow.NewColumn(schema.Field(i+1), &floatChunks[i]))
	}
	for i := 0; i < 1000; i++ {
		columns = append(columns, *arrow.NewColumn(schema.Field(i+1001), &boolChunks[i]))
	}

	// 创建 Arrow Table
	table := array.NewTable(schema, columns, int64(len(turbineData)))
	return &table, nil
}

// writeParquet 写入 Parquet 文件
func writeParquet(table *arrow.Table, filename string) error {
	// 创建 Parquet 文件
	f, err := os.Create("../../" + filename)
	if err != nil {
		return err
	}
	defer f.Close()

	// 写入 Parquet
	err = pqarrow.WriteTable(*table, f, 1024, nil, pqarrow.DefaultWriterProps())
	return err
}

// 记录所有风机数据的内存存储
var windTurbineData = make(map[string][]*pb.WindTurbineData)
var mu sync.Mutex

// 处理数据写入
func (s *WindTurbineServer) SendData(ctx context.Context, data *pb.WindTurbineData) (*pb.WriteResponse, error) {
	mu.Lock()
	defer mu.Unlock()

	// 将新数据追加到对应风机的数组中
	windTurbineData[data.TurbineID] = append(windTurbineData[data.TurbineID], data)

	// 检查是否达到 1800 条数据
	if len(windTurbineData[data.TurbineID]) >= 1800 {

		// 调用 createArrowTable 生成 Arrow Table
		start1 := time.Now()
		table, err := createArrowTable(windTurbineData[data.TurbineID])
		// 打印createArrowTable耗时
		fmt.Printf("turbine: %screateArrowTable耗时：%v\n", data.TurbineID, time.Since(start1))
		if err != nil {
			return nil, fmt.Errorf("failed to create Arrow table: %v", err)
		}

		// 写入 Parquet 文件
		start2 := time.Now()
		// parquetFIle 加windTurbineData[data.TurbineID][0].Timestamp前缀
		parquetFilename := fmt.Sprintf("%s_%d.parquet", data.TurbineID, windTurbineData[data.TurbineID][0].Timestamp)

		err = writeParquet(table, parquetFilename)
		fmt.Printf("turbine: %swriteParquet耗时：%v\n", data.TurbineID, time.Since(start2))
		if err != nil {
			return nil, fmt.Errorf("failed to write Parquet file: %v", err)
		}

		// 清空该风机的数据
		delete(windTurbineData, data.TurbineID)

	}

	// 返回响应
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
