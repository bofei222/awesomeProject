package main

import (
	"fmt"
	"math/rand"
	"time"
)

// 定义数据结构
var boolData = make([][][]bool, 100)     // 存储布尔值数据
var floatData = make([][][]float32, 100) // 存储浮点数数据

func init() {
	// 初始化所有数组，每台风机有 3600 秒数据
	for i := 0; i < 100; i++ {

		boolData[i] = make([][]bool, 1000)     // 每台风机 1000 个布尔变量
		floatData[i] = make([][]float32, 1000) // 每台风机 1000 个浮点数变量

		for j := 0; j < 1000; j++ {
			boolData[i][j] = make([]bool, 3600)
			floatData[i][j] = make([]float32, 3600)
		}
	}
}

// 模拟生成每台风机每秒的数据
func generateData() (int64, []bool, []float32) {
	timestamp := time.Now().UnixMilli()

	// 生成 1000 个布尔变量
	bools := make([]bool, 1000)
	for i := range bools {
		bools[i] = rand.Intn(2) == 1
	}

	// 生成 1000 个浮点数变量
	floats := make([]float32, 1000)
	for i := range floats {
		floats[i] = rand.Float32() * 100
	}

	return timestamp, bools, floats
}

// 将数据写入到数组中
func writeData(turbineID, second int, timestamp int64, bools []bool, floats []float32) {

	for i := 0; i < 1000; i++ {
		boolData[turbineID][i][second] = bools[i]
		floatData[turbineID][i][second] = floats[i]
	}
}

func main() {
	// 模拟每台风机每秒写入一条数据，持续 3600 秒
	for second := 0; second < 3600; second++ {
		for turbineID := 0; turbineID < 100; turbineID++ {
			timestamp, bools, floats := generateData()
			writeData(turbineID, second, timestamp, bools, floats)
		}

	}

	// 输出某台风机在第 0 秒的数据，检查是否写入成功
	fmt.Println("Example Data for Turbine 0 at second 0:")

	for i := 0; i < 5; i++ {
		fmt.Printf("Bool_var_%d: %v\n", i+1, boolData[0][i][0])
		fmt.Printf("Float_var_%d: %v\n", i+1, floatData[0][i][0])
	}
}
