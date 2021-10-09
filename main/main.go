package main

import (
	"awesomeProject/kafka"
	"fmt"
	"os"
)

func main() {
	//client := kafka.NewKafkaClient()
	kafka.Init()
	client := kafka.NewKafkaClient()
	//client.RecMsg2(func(message *sarama.ConsumerMessage) {
	//	fmt.Println(message)
	//})

	file, err := os.ReadFile("C:\\Users\\bofei\\Desktop\\11\\event.dat")
	if nil != err {
		fmt.Println(err)
	}
	client.SyncSend("windPower", string(file))

}
