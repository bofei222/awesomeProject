package mystring

import (
	"fmt"
	"strings"
	"testing"
)

func TestString(t *testing.T) {
	startFaults := []string{
		"5ff",
		"6ff",
	}
	for i, startFault := range startFaults {
		if strings.Contains(startFault, "ff") {
			// 改变 这个副本是不行的
			startFault = strings.ReplaceAll(startFault, "ff", "")
			fmt.Println(startFault)
			startFaults[i] = startFault // 将数组内 保存的 引用改变
		}
	}
	fmt.Println(startFaults)
}
