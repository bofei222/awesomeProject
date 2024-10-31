package main

import (
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"runtime"
	"time"
)

// WindTurbineData 结构体，存储单台风机的列式数据
type WindTurbineData struct {
	BoolCols  map[string][]bool    // MA 开头的列
	FloatCols map[string][]float32 // MC 开头的列
	//Timestamps []int64              // 时间戳（毫秒数）
}

// FarmData 存储整个风场数据
type FarmData struct {
	Turbines map[int]*WindTurbineData // 多台风机的数据
}

// NewWindTurbineData 初始化单台风机的数据结构
func NewWindTurbineData() *WindTurbineData {
	return &WindTurbineData{
		BoolCols:  make(map[string][]bool, 1000),
		FloatCols: make(map[string][]float32, 1000),
		//Timestamps: make([]int64, 0, 3600),
	}
}

// NewFarmData 初始化整个风场数据
func NewFarmData(numTurbines int) *FarmData {
	farm := &FarmData{Turbines: make(map[int]*WindTurbineData, numTurbines)}
	for i := 0; i < numTurbines; i++ {
		farm.Turbines[i] = NewWindTurbineData()
	}
	return farm
}

// 添加一秒的数据到指定风机
func (data *WindTurbineData) AddRow(timestamp int64, bools []bool, floats []float32) {
	//data.Timestamps = append(data.Timestamps, timestamp)

	for i, val := range bools {
		key := fmt.Sprintf("MA%03d", i+1)
		data.BoolCols[key] = append(data.BoolCols[key], val)
	}

	for i, val := range floats {
		key := fmt.Sprintf("MC%03d", i+1)
		data.FloatCols[key] = append(data.FloatCols[key], val)
	}
}

// 生成随机数据
func generateRandomData() ([]bool, []float32) {
	bools := make([]bool, 1000)
	floats := make([]float32, 1000)

	for i := range bools {
		bools[i] = rand.Intn(2) == 1
	}

	for i := range floats {
		floats[i] = rand.Float32() * 100.0
	}

	return bools, floats
}

// 打印内存使用情况
func printMemoryUsage() {
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	fmt.Printf("Alloc = %v MiB", m.Alloc/1024/1024)
	fmt.Printf("\tTotalAlloc = %v MiB", m.TotalAlloc/1024/1024)
	fmt.Printf("\tSys = %v MiB", m.Sys/1024/1024)
	fmt.Printf("\tNumGC = %v\n", m.NumGC)
}

func main() {
	// 启动 pprof
	go func() {
		log.Println(http.ListenAndServe("localhost:6060", nil))
	}()

	farm := NewFarmData(100) // 初始化 100 台风机

	startTime := time.Now().UnixMilli() // 起始时间戳

	// 为每台风机生成 3600 秒的数据
	for i := 0; i < 3600; i++ {
		timestamp := startTime + int64(i*1000)
		bools, floats := generateRandomData()

		for _, turbine := range farm.Turbines {
			turbine.AddRow(timestamp, bools, floats)
		}
	}

	fmt.Println("数据填充完成。")
	printMemoryUsage()

	// 查询示例：风机 0 的 MC001 的平均值
	mean, err := farm.Turbines[0].GetFloatColumnMean("MC001")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("风机 0 的 MC001 平均值: %.2f\n", mean)

	// 计算全场风场的MC001的平均值

	printMemoryUsage()
}

// 获取浮点列的平均值
func (data *WindTurbineData) GetFloatColumnMean(colName string) (float32, error) {
	column, exists := data.FloatCols[colName]
	if !exists {
		return 0, fmt.Errorf("column %s not found", colName)
	}

	var sum float32
	for _, val := range column {
		sum += val
	}
	return sum / float32(len(column)), nil
}
