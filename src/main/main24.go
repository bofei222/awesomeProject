package main

import (
	"fmt"
	"time"
)

// SplitTimeRange 按整点切割时间段，支持跨天和时区
func SplitTimeRange(start, end time.Time) []time.Time {
	var times []time.Time

	// 计算开始时间的下一个整点时间
	current := time.Date(start.Year(), start.Month(), start.Day(), start.Hour()+1, 0, 0, 0, start.Location())
	for current.Before(end) {
		times = append(times, current)
		current = current.Add(time.Hour)
	}

	return times
}

func main() {
	// 定义时区
	loc, err := time.LoadLocation("Asia/Shanghai")
	if err != nil {
		fmt.Println("Error loading location:", err)
		return
	}

	// 定义起始时间和结束时间
	start := time.Date(2023, 7, 28, 23, 30, 0, 0, loc)
	end := time.Date(2023, 7, 29, 2, 30, 0, 0, loc)

	// 获取切割后的时间点
	times := SplitTimeRange(start, end)

	// 打印结果
	for _, t := range times {
		fmt.Println(t)
	}
}
