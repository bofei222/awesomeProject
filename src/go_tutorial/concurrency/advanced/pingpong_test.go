package advanced

import (
	"fmt"
	"testing"
	"time"
)

func TestPingPong(t *testing.T) {
	table := make(chan *Ball)
	go player("ping", table)
	go player("pong", table)
	table <- new(Ball)
	time.Sleep(1 * time.Second)
	//<- table
}

func player(name string, table chan *Ball) {
	for {
		ball := <-table
		ball.hists++
		fmt.Println(name, ball.hists)
		time.Sleep(100 * time.Millisecond)
		table <- ball
	}
}

type Ball struct {
	hists int
}
