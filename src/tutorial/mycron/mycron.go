package main

import (
	"github.com/robfig/cron/v3"
	"log"
	"time"
)

func main() {
	//c := cron.New(cron.WithSeconds())

	//c := cron.New(cron.WithSeconds(), cron.WithChain(cron.SkipIfStillRunning(cron.VerbosePrintfLogger(log.New(os.Stdout, "mycron: ", log.LstdFlags)))))

	c := cron.New(cron.WithSeconds())
	//spec := "*/5 * * * * *" // 每隔5s执行一次，cron格式（秒，分，时，天，月，周）

	//// 添加一个任务
	//c.AddFunc(spec, func() {
	//	log.Printf("111 time = %d\n", time.Now().Unix())
	//})

	// 添加一个任务
	c.AddFunc("*/1 * * * * *", func() { // 可以随时添加多个定时任务
		log.Printf("111")
		time.Sleep(time.Second * 3)
	})

	// 不仅本任务的 下次执行会被跳过 ， 其他 任务的 也不会执行
	c.AddFunc("*/1 * * * * *", func() { // 可以随时添加多个定时任务
		log.Printf("2222")
	})
	c.Start()
	select {}
}
