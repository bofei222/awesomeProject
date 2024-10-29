package main

import (
	"fmt"
	"log"
)

// WindTurbineData 结构体，存储列式数据
type WindTurbineData struct {
	BoolCols   map[string][]bool    // MA 开头的列
	IntCols    map[string][]int32   // MB 开头的列
	FloatCols  map[string][]float32 // MC 开头的列
	Timestamps []string             // 时间戳
}

// 初始化 WindTurbineData
func NewWindTurbineData() *WindTurbineData {
	return &WindTurbineData{
		BoolCols:   make(map[string][]bool),
		IntCols:    make(map[string][]int32),
		FloatCols:  make(map[string][]float32),
		Timestamps: []string{},
	}
}

// 向数据结构中插入一行数据
func (data *WindTurbineData) AddRow(timestamp string, bools []bool, ints []int32, floats []float32) {
	data.Timestamps = append(data.Timestamps, timestamp)

	// 插入布尔型数据
	for i, val := range bools {
		key := fmt.Sprintf("MA%03d", i+1)
		data.BoolCols[key] = append(data.BoolCols[key], val)
	}

	// 插入整型数据
	for i, val := range ints {
		key := fmt.Sprintf("MB%03d", i+1)
		data.IntCols[key] = append(data.IntCols[key], val)
	}

	// 插入浮点型数据
	for i, val := range floats {
		key := fmt.Sprintf("MC%03d", i+1)
		data.FloatCols[key] = append(data.FloatCols[key], val)
	}
}

// 计算 MC 列的平均值
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

// 计算 MB 列的求和
func (data *WindTurbineData) GetIntColumnSum(colName string) (int32, error) {
	column, exists := data.IntCols[colName]
	if !exists {
		return 0, fmt.Errorf("column %s not found", colName)
	}

	var sum int32
	for _, val := range column {
		sum += val
	}
	return sum, nil
}

func main() {
	data := NewWindTurbineData()

	// 模拟插入三行数据
	data.AddRow("2024/10/29 11:49:01",
		[]bool{true, true, true},
		[]int32{6, 6, 6},
		[]float32{11.22, 11.22, 11.22})

	data.AddRow("2024/10/30 11:49:02",
		[]bool{true, true, true},
		[]int32{7, 7, 7},
		[]float32{33.44, 33.44, 33.44})

	data.AddRow("2024/10/31 11:49:03",
		[]bool{true, true, true},
		[]int32{8, 8, 8},
		[]float32{55.66, 55.66, 55.66})

	// 查询 MC001 的平均值
	mean, err := data.GetFloatColumnMean("MC001")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("MC001 Mean: %.2f\n", mean)

	// 查询 MB001 的求和
	sum, err := data.GetIntColumnSum("MB001")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("MB001 Sum: %d\n", sum)
}
