package main

import (
	"fmt"
	"github.com/tealeg/xlsx"
)

func main() {
	// 打开 Excel 文件
	xlFile, err := xlsx.OpenFile("D:\\下载\\DBCFCA_pub_20240306-094638\\20240306-094638\\pubs_20240306-094639.xlsx")
	if err != nil {
		fmt.Println("无法打开 Excel 文件:", err)
		return
	}

	// 获取第四个工作表
	sheetIndex := 4 // 第四个工作表的索引是3（索引从0开始）
	sheet := xlFile.Sheets[sheetIndex]

	// 定义要读取的行和列范围
	startRow := 19 // 20 行在数组索引中是 19
	endRow := 1074 // 1075 行在数组索引中是 1074

	startCol := 52 // 53 列在数组索引中是 52 (BA 列)

	// 定义存储数据的二维数组
	var data [][]string

	// 读取数据并存储到数组中
	var rowData []string
	for rowIndex := startRow; rowIndex <= endRow; rowIndex++ {
		row := sheet.Rows[rowIndex]
		cell := row.Cells[startCol]
		text := cell.String()
		rowData = append(rowData, text)

		if (rowIndex+1-startRow)%32 == 0 { // 每32行为一组数据
			data = append(data, rowData)
			rowData = nil // 重置rowData
		}
	}

	// 输出每个数组的数据
	for _, group := range data {
		//fmt.Printf("Group %d:\n", i+1)

		for i := len(group) - 1; i >= 0; i-- {
			fmt.Println(group[i])
		}

	}
}
