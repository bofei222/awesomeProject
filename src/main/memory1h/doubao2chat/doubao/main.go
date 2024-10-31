package main

import (
	"fmt"
	"time"
)

// DataPoint 表示单个数据点，包含1000个float32和1000个bool类型字段以及一个时间戳
type DataPoint struct {
	Floats    [1000]float32
	Bools     [1000]bool
	Timestamp int64
}

// WindTurbine 表示单个风机的数据，包含3600秒的数据点
type WindTurbine struct {
	DataPoints [3600]DataPoint
}

// SimulateData 用于模拟所有风机的数据
func SimulateData() []WindTurbine {
	var turbines []WindTurbine
	startTime := time.Now().Unix()

	for i := 0; i < 100; i++ {
		turbine := WindTurbine{}
		for j := 0; j < 3600; j++ {
			dataPoint := DataPoint{
				Timestamp: startTime + int64(j),
			}
			// 模拟float32数据
			for k := range dataPoint.Floats {
				dataPoint.Floats[k] = float32仿真数据(k)
			}
			// 模拟bool数据
			for k := range dataPoint.Bools {
				dataPoint.Bools[k] = 仿真布尔数据(k)
			}
			turbine.DataPoints[j] = dataPoint
		}
		turbines = append(turbines, turbine)
	}

	return turbines
}

// 仿真float32数据的函数，这里只是一个示例，实际中应根据需要生成数据
func float32仿真数据(index int) float32 {
	// 这里只是返回了一个简单的模拟值，实际应用中应根据具体需求生成数据
	return float32(index) * 0.1
}

// 仿真布尔数据的函数，这里只是一个示例，实际中应根据需要生成数据
func 仿真布尔数据(index int) bool {
	// 这里只是返回了一个简单的模拟值，实际应用中应根据具体需求生成数据
	return index%2 == 0
}

func main() {
	// 模拟数据并存储到内存中
	windTurbineData := SimulateData()

	// 打印第一台风机的第一个数据点作为示例
	if len(windTurbineData) > 0 {
		fmt.Printf("First turbine first data point: %+v\n", windTurbineData[0].DataPoints[0])
	}
}
