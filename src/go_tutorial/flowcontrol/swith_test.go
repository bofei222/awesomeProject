package flowcontrol

import (
	"fmt"
	"runtime"
	"testing"
	"time"
)

func TestSwitch(t *testing.T) {
	fmt.Print("Go runs on ")
	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("OS X.")
	case "linux":
		fmt.Println("Linux.")
	default:
		// freebsd, openbsd,
		// plan9, windows...
		fmt.Printf("%s.\n", os)
	}
}
func TestSwitchCaseOrder(t *testing.T) {
	fmt.Println("When's Saturday?")
	today := time.Now().Weekday()
	switch time.Saturday {
	case today + 0:
		fmt.Println("Today.")
	case today + 1:
		fmt.Println("Tomorrow.")
	case today + 4:
		fmt.Println("In 4 days.")
	default:
		fmt.Println("Too far away.")
	}
}

func TestNotConditionSwitch(t *testing.T) {
	time := time.Now()
	switch {
	case time.Hour() < 12:
		fmt.Println("Good morning!")
	case time.Hour() < 17:
		fmt.Println("Good afternoon.")
	default:
		fmt.Println("Good evening.")
	}
}
