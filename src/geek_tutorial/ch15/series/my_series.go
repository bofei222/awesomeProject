package series

import "errors"

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
