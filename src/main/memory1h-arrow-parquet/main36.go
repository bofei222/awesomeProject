package main

import (
	"fmt"
	"math/rand"
	"net/http"
	_ "net/http/pprof"
	"os"
	"time"
	"unsafe"

	"github.com/apache/arrow-go/v18/arrow"
	"github.com/apache/arrow-go/v18/arrow/array"
	"github.com/apache/arrow-go/v18/arrow/memory"
	"github.com/apache/arrow-go/v18/parquet/pqarrow"
)

// WindTurbineData 定义了单条数据的结构
type WindTurbineData struct {
	Timestamp int64         // 时间戳，粒度为秒
	FloatData [1000]float32 // 1000 个 float32 点位
	BoolData  [1000]bool    // 1000 个 bool 点位
}

// simulateWindTurbineData 模拟单个风机的数据
func simulateWindTurbineData(startTime time.Time, durationInSeconds int) []WindTurbineData {
	data := make([]WindTurbineData, 0, durationInSeconds)

	for i := 0; i < durationInSeconds; i++ {
		timestamp := startTime.Add(time.Duration(i) * time.Second).Unix()

		// 随机生成 1000 个 float32 类型点位和 1000 个 bool 类型点位
		var floatData [1000]float32
		var boolData [1000]bool
		for j := 0; j < 1000; j++ {
			floatData[j] = rand.Float32() * 100.0 // 随机浮点数，范围在 0 到 100 之间
			boolData[j] = rand.Intn(2) == 1       // 随机生成 true 或 false
		}

		data = append(data, WindTurbineData{
			Timestamp: timestamp,
			FloatData: floatData,
			BoolData:  boolData,
		})
	}

	return data
}

// createArrowTable 创建 Arrow Table，将每台风机的数据转换为 Arrow 格式
func createArrowTable(turbineData []WindTurbineData) (*arrow.Table, error) {
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
	chunks := []arrow.Array{
		timestampBuilder.NewArray(),
	}

	for i := 0; i < 1000; i++ {
		chunks = append(chunks, floatBuilders[i].NewArray())
		chunks = append(chunks, boolBuilders[i].NewArray())
	}

	// 创建 Arrow Table
	table := array.NewTable(schema, chunks, int64(len(turbineData)))
	return table, nil
}

// writeParquet 写入 Parquet 文件
func writeParquet(table *arrow.Table, filename string) error {
	// 创建 Parquet 文件
	f, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer f.Close()

	// 写入 Parquet
	err = pqarrow.WriteTable(*table, f, 1024, nil, pqarrow.DefaultWriterProps())
	return err
}

func main() {
	// 启动 pprof HTTP 服务器
	go func() {
		fmt.Println("Starting pprof server at :6060")
		http.ListenAndServe("localhost:6060", nil)
	}()

	// 模拟 120 台风机，3600 秒的数据
	numWindTurbines := 120
	durationInSeconds := 3600
	startTime := time.Now()

	// 存储所有风机的数据
	allWindTurbinesData := make([][]WindTurbineData, numWindTurbines)

	// 模拟每台风机的数据
	for i := 0; i < numWindTurbines; i++ {
		allWindTurbinesData[i] = simulateWindTurbineData(startTime, durationInSeconds)
	}

	// 打印一些示例数据
	for i := 0; i < numWindTurbines; i++ {
		fmt.Printf("风机 %d 的第 1 秒数据: %+v\n", i+1, allWindTurbinesData[i][0])
	}

	// 打印 map 的内存大小
	fmt.Printf("Size of map: %d bytes\n", unsafe.Sizeof(allWindTurbinesData))
	fmt.Println("数据生成完毕！")

	// 每台风机写入一个 Parquet 文件
	for i := 0; i < numWindTurbines; i++ {
		table, err := createArrowTable(allWindTurbinesData[i])
		if err != nil {
			fmt.Printf("Error creating Arrow table for wind turbine %d: %v\n", i+1, err)
			continue
		}

		// 写入 Parquet 文件
		parquetFilename := fmt.Sprintf("wind_turbine_%d.parquet", i+1)
		err = writeParquet(table, parquetFilename)
		if err != nil {
			fmt.Printf("Error writing Parquet file for wind turbine %d: %v\n", i+1, err)
		} else {
			fmt.Printf("Successfully wrote Parquet file for wind turbine %d: %s\n", i+1, parquetFilename)
		}
	}

	// 等待一段时间，以便在 pprof 服务器上查看数据
	time.Sleep(99999999 * time.Minute)
}
