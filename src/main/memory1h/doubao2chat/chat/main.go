package main

import (
	"fmt"
	"math/rand"
	"time"
)

// 定义风机数据结构
type WindTurbineData struct {
	Timestamp int64     // 时间戳（秒级）
	Floats    []float32 // 1000 个 float32 类型数据
	Bools     []bool    // 1000 个 bool 类型数据
}

// 模拟数据生成
func generateTurbineData(startTime int64) []*WindTurbineData {
	var data []*WindTurbineData
	for i := int64(0); i < 3600; i++ {
		timestamp := startTime + i
		floatData := make([]float32, 1000)
		boolData := make([]bool, 1000)

		// 随机生成 float32 和 bool 数据
		for j := 0; j < 1000; j++ {
			floatData[j] = float32(rand.Float32()) // 使用随机浮点数作为示例
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

func main() {
	rand.Seed(time.Now().UnixNano()) // 设置随机种子

	// 用 map 存储 10 台风机的数据
	windTurbines := make(map[int][]*WindTurbineData)

	// 开始时间戳（秒级，模拟 3600 秒）
	startTime := time.Now().Unix()

	// 模拟每台风机的数据
	for i := 0; i < 100; i++ {
		windTurbines[i] = generateTurbineData(startTime)
		fmt.Printf("风机 %d 数据生成完成\n", i+1)
	}

	// 打印示例数据
	for i, turbineData := range windTurbines {
		fmt.Printf("风机 %d 第一条数据: %+v\n", i+1, turbineData[0])
	}
}
