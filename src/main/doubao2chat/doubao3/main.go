package main

import (
	"fmt"
	"math/rand"
	"time"
)

// DataPoint 表示单个数据点，包含1000个float32和1000个bool类型字段以及一个时间戳
type DataPoint struct {
	Timestamp int64
	Floats    [1000]float32
	Bools     [1000]bool
}

// SimulateData 用于模拟单台风机的数据
func SimulateData() []*DataPoint {
	var dataPoints []*DataPoint
	for j := 0; j < 3600; j++ {
		dataPoint := &DataPoint{
			Timestamp: int64(j) + time.Now().Unix(),
		}
		for k := range dataPoint.Bools {
			// 模拟bool数据
			// 模拟float32数据
			dataPoint.Floats[k] = float32(rand.Float32()) // 使用随机浮点数作为示例
			dataPoint.Bools[k] = rand.Intn(2) == 0        // 随机生成true或false
		}
		dataPoints = append(dataPoints, dataPoint)
	}
	return dataPoints
}

func main() {
	// 存储所有风机的数据
	allTurbinesData := make([][]*DataPoint, 100)

	// 模拟10台风机的数据
	for i := 0; i < 100; i++ {
		allTurbinesData[i] = SimulateData()
	}

	// 打印第一台风机的第一个数据点作为示例
	if len(allTurbinesData) > 0 && len(allTurbinesData[0]) > 0 {
		fmt.Printf("First turbine first data point: %+v\n", allTurbinesData[0][0])
	}
}
