package error

import (
	"errors"
	"testing"
)

// 两种错误 ，让调用者区分
// var定义两种错误 var AError error = errors.New
// 尽早失败
// 尽早return continue 减少嵌套
func GetFibonacci(n int) ([]int, error) {
	if n < 1 {
		return nil, errors.New("error")
	}
	fibList := []int{1, 1}

	// 最新的一位数 总是 等于 次次新 + 次新
	for i := 2; i < n; i++ {
		fibList = append(fibList, fibList[i-2]+fibList[i-1])
	}
	return fibList, nil
}

func TestGetFibonacci(t *testing.T) {
	//t.Log(GetFibonacci(-10))

	// 两段表达式
	if fibonacci, err := GetFibonacci(-10); err != nil {
		t.Log(err)
	} else {
		t.Log(fibonacci)
	}

}
