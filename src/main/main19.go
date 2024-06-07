package main

import (
	"fmt"
	"strconv"
	"time"
)

func makeCakeAndSend(cs chan string, count int) {
	for i := 1; i <= count; i++ {
		cakeName := "Strawberry Cake " + strconv.Itoa(i)
		// 睡眠1s
		time.Sleep(1 * 1e9) //sleep for 1 second
		cs <- cakeName      //send a strawberry cake
	}
}

func receiveCakeAndPack(cs chan string) {
	for s := range cs {
		fmt.Println("Packing received cake: ", s)
	}
}

func main() {
	cs := make(chan string)
	go makeCakeAndSend(cs, 5)
	go receiveCakeAndPack(cs)

	//sleep for a while so that the program doesn’t exit immediately
	time.Sleep(10 * 1e9)
}
