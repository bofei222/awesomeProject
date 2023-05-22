package main

import (
	"fmt"
	"time"
)

func main() {
	var t int64
	t = 1684740483082
	format := time.UnixMilli(t).Format("2006-01-02T15:04:05+08:00")

	parse, err := time.Parse(time.RFC3339, format)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(parse)
	}

	fmt.Println(time.UnixMilli(t))

}
