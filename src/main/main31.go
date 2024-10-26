package main

import (
	"fmt"
	"math/rand"
	"net/http"
	_ "net/http/pprof" // 导入 pprof 包以启用 pprof 路由
	"time"
)

func main() {
	// 启动 pprof HTTP 服务器
	go func() {
		fmt.Println("Starting pprof server at :6060")
		http.ListenAndServe("localhost:6060", nil)
	}()

	rand.Seed(time.Now().UnixNano())

	const turbineCount = 10
	const boolCount = 1000
	const floatCount = 1000
	const dataPoints = 3600

	data := make(map[string][]map[string]interface{})

	// 生成数据
	//
	start := time.Now()
	funcName(turbineCount, dataPoints, boolCount, floatCount, data)
	// 生成数据耗时：
	fmt.Println("生成数据耗时：", time.Since(start))
	// 每30秒查询一次0003风机的MC004~MC024的平均值
	ticker30s := time.NewTicker(30 * time.Second)
	go func() {
		for range ticker30s.C {
			start := time.Now()
			avgValues := calculateAverage(data["003"], "MC004", "MC024")
			// 计算calculateAverage耗时并打印
			fmt.Println("calculateAverage耗时：")

			fmt.Printf("Average values :%vfor turbine 0003 (MC004 - MC024): %v\n", time.Since(start), avgValues)
		}
	}()

	// 每3分钟查询一次全部风机的MC004~MC024的平均值的平均
	ticker3m := time.NewTicker(3 * time.Minute)
	go func() {
		for range ticker3m.C {
			totalAvgValues := make([]float32, 21) // MC004 to MC024
			count := 0

			for i := 0; i < turbineCount; i++ {
				turbineNum := fmt.Sprintf("%03d", i+1)
				avgValues := calculateAverage(data[turbineNum], "MC004", "MC024")
				for j := 0; j < len(avgValues); j++ {
					totalAvgValues[j] += avgValues[j]
				}
				count++
			}

			for i := range totalAvgValues {
				totalAvgValues[i] /= float32(count)
			}

			fmt.Printf("Overall average values for all turbines (MC004 - MC024): %v\n", totalAvgValues)
		}
	}()

	// 让程序运行一段时间
	time.Sleep(60 * time.Minute)
	ticker30s.Stop()
	ticker3m.Stop()
}

func funcName(turbineCount int, dataPoints int, boolCount int, floatCount int, data map[string][]map[string]interface{}) {
	for i := 0; i < turbineCount; i++ {
		turbineNum := fmt.Sprintf("%03d", i+1)
		data[turbineNum] = make([]map[string]interface{}, 0, dataPoints) // 初始化一个容量为3600的slice
		for j := 0; j < dataPoints; j++ {
			// map的大小boolCount+floatCount+1
			point := make(map[string]interface{}, boolCount+floatCount+1)
			point["timestamp"] = time.Now().Add(time.Duration(j) * time.Second).Unix()

			for k := 1; k <= boolCount; k++ {
				key := fmt.Sprintf("MA%03d", k)
				point[key] = rand.Intn(2) == 1
			}

			for k := 1; k <= floatCount; k++ {
				key := fmt.Sprintf("MC%03d", k)
				point[key] = rand.Float32() * 100
			}

			data[turbineNum] = append(data[turbineNum], point)
		}
		fmt.Println(len(data))
	}
}

// calculateAverage calculates the average values for the specified range of keys in the data.
func calculateAverage(points []map[string]interface{}, startKey, endKey string) []float32 {
	startIndex := getIndex(startKey)
	endIndex := getIndex(endKey)
	count := float32(len(points))
	averages := make([]float32, endIndex-startIndex+1)

	for _, point := range points {
		for i := startIndex; i <= endIndex; i++ {
			key := fmt.Sprintf("MC%03d", i)
			if value, ok := point[key]; ok {
				averages[i-startIndex] += value.(float32)
			}
		}
	}

	for i := range averages {
		if count > 0 {
			averages[i] /= count
		}
	}

	return averages
}

// getIndex retrieves the index for MC keys.
func getIndex(key string) int {
	var index int
	fmt.Sscanf(key, "MC%03d", &index)
	return index
}
