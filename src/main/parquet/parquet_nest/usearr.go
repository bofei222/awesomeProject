package main

import (
	"fmt"
	"math/rand"
	"net/http"
	"time"
)

// 定义风机数据结构
type WindTurbineData struct {
	Timestamp int64     // 时间戳（秒级）
	Floats    []float32 // 1000 个 float32 类型数据
	Bools     []bool    // 1000 个 bool 类型数据
}

//	MC001 :0 ,MC2000: 1999
//
// 模拟数据生成
func generateTurbineData(startTime int64) []*WindTurbineData {
	var data []*WindTurbineData
	for i := int64(0); i < 3600; i++ {
		timestamp := startTime + i
		floatData := make([]float32, 1000)
		boolData := make([]bool, 1000)

		// 随机生成 float32 和 bool 数据
		for j := 0; j < 1000; j++ {
			floatData[j] = rand.Float32() * 100.0 // 假设 float 数据范围 0-100
			boolData[j] = rand.Intn(2) == 1
		}

		data = append(data, &WindTurbineData{
			Timestamp: timestamp,
			Floats:    floatData,
			Bools:     boolData,
		})
	}
	return data
}

// 计算风机 float 数组中第 4 到第 24 个元素的 3600 秒平均值
func calculateAverage(turbineData []*WindTurbineData) []float32 {
	averages := make([]float32, 21)
	for i := 3; i <= 23; i++ { // 索引从 3 到 23，对应第 4 到第 24 个元素
		var sum float32
		for _, data := range turbineData {
			sum += data.Floats[i]
		}
		averages[i-3] = sum / 3600
	}
	return averages
}

// 计算全场风机 float 数组中第 4 到第 24 个元素的平均值
func calculateAllTurbinesAverage(windTurbines map[int][]*WindTurbineData) []float32 {
	allAverages := make([]float32, 21)
	for i := 0; i < 21; i++ {
		var totalSum float32
		for _, turbineData := range windTurbines {
			avg := calculateAverage(turbineData)
			totalSum += avg[i]
		}
		allAverages[i] = totalSum / float32(len(windTurbines))
	}
	return allAverages
}

func main() {
	// 启动 pprof HTTP 服务器
	go func() {
		fmt.Println("Starting pprof server at :6060")
		http.ListenAndServe("localhost:6060", nil)
	}()

	rand.Seed(time.Now().UnixNano()) // 设置随机种子

	// 用 map 存储 10 台风机的数据
	windTurbines := make(map[int][]*WindTurbineData)

	// 开始时间戳（秒级，模拟 3600 秒）
	startTime := time.Now().Unix()

	// 模拟每台风机的数据
	for i := 0; i < 1; i++ {
		windTurbines[i] = generateTurbineData(startTime)
		fmt.Printf("风机 %04d 数据生成完成\n", i)
	}

}
