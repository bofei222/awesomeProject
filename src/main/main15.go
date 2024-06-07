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

	startCol := 49 // 50 列在数组索引中是 49 (AX 列)
	endCol := 54   // 56 列在数组索引中是 55 (BC 列)

	// 遍历指定行和列范围内的单元格
	for rowIndex := startRow; rowIndex <= endRow; rowIndex++ {
		row := sheet.Rows[rowIndex]
		for colIndex := startCol; colIndex <= endCol; colIndex++ {
			cell := row.Cells[colIndex]
			text := cell.String()
			fmt.Printf("%s\t", text)
		}
		fmt.Println()
	}
}
