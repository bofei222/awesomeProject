package main

import (
	"fmt"
	"math/rand"
	"net/http"
	_ "net/http/pprof" // 导入 pprof 包以启用 pprof 路由
	"time"
	"unsafe"
)

// 定义单条数据的结构
type WindTurbineData struct {
	Timestamp int64         // 时间戳，粒度为秒
	FloatData [1000]float32 // 1000 个 float32 点位
	BoolData  [1000]bool    // 1000 个 bool 点位
}

// 模拟单个风机的数据
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
	// 等待一段时间，以便在 pprof 服务器上查看数据
	time.Sleep(99999999 * time.Minute)
}
