package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	// 指定文件路径
	filePath := "D:/bofeiProjects/sany/WindSmartCCM/event-center/example/config/config.yaml"

	// 使用 ioutil.ReadFile 读取整个文件的内容
	content, err := ioutil.ReadFile(filePath)
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	// 打印文件内容
	fmt.Println("File Content:")
	fmt.Println(string(content))

	// 另一种方法：使用 os.Open 和 bufio.Scanner 按行读取文件
	file, err := os.Open(filePath)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close() // 确保文件在函数退出时被关闭

	fmt.Println("File Content Line by Line:")
	// 使用 bufio.Scanner 逐行读取文件内容
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}

	// 检查扫描时是否发生错误
	if err := scanner.Err(); err != nil {
		fmt.Println("Error scanning file:", err)
	}
}
