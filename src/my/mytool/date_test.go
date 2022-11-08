package mytool

import (
	"fmt"
	testing2 "testing"
	"time"
)

func TestDatePlus(t *testing2.T) {

	now := time.Now()
	fmt.Println(now)

	date := now.Format("20060102")
	fmt.Println(date)
	date = now.Format("2006-01-02")
	fmt.Println(date)

	date2, err := time.Parse("2006-01-02", "2021-01-01")
	if err == nil {
		fmt.Println(date2)
	}
	addDate := date2.AddDate(0, 0, 1)
	result := addDate.Format("2006-01-02")
	fmt.Println(result)
}
