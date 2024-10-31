package main

import (
	"fmt"
	"math/rand"
	"time"
)

// DataPoint 表示单个数据点，包含1000个float32和1000个bool类型字段以及一个时间戳
type DataPoint struct {
	Floats    [1000]float32
	Bools     [1000]bool
	Timestamp int64
}

// 模拟float32数据的函数
func generateFloatData() [1000]float32 {
	var floats [1000]float32
	for i := range floats {
		floats[i] = float32(rand.Float32()) // 使用随机浮点数作为示例
	}
	return floats
}

// 模拟bool数据的函数
func generateBoolData() [1000]bool {
	var bools [1000]bool
	for i := range bools {
		bools[i] = rand.Intn(2) == 0 // 随机生成true或false
	}
	return bools
}

// SimulateData 用于模拟所有风机的数据
func SimulateData() [][]DataPoint {
	var allTurbinesData [][]DataPoint
	for i := 0; i < 100; i++ {
		var turbineData []DataPoint
		for j := 0; j < 3600; j++ {
			dataPoint := DataPoint{
				Floats:    generateFloatData(),
				Bools:     generateBoolData(),
				Timestamp: int64(j) + time.Now().Unix(),
			}
			turbineData = append(turbineData, dataPoint)
		}
		allTurbinesData = append(allTurbinesData, turbineData)
	}
	return allTurbinesData
}

func main() {
	// 模拟数据并存储到内存中
	windTurbineData := SimulateData()

	// 打印第一台风机的第一个数据点作为示例
	if len(windTurbineData) > 0 && len(windTurbineData[0]) > 0 {
		fmt.Printf("First turbine first data point: %+v\n", windTurbineData[0][0])
	}
}
