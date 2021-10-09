package main

import "fmt"
import "github.com/robfig/cron"

func main() {
	c := cron.New() //精确到秒

	//定时任务
	spec := "*/1 * * * * ?" //cron表达式，每秒一次
	c.AddFunc(spec, func() { fmt.Println("11111") })
	c.AddFunc(spec, func() { fmt.Println("11111") })

	c.Start()
	select {} //阻塞主线程停止
}
