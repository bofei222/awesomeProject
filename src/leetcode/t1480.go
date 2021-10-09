package main

import "fmt"

//输入：nums = [1,2,3,4]
//输出：[1,3,6,10]
//解释：动态和计算过程为 [1, 1+2, 1+2+3, 1+2+3+4] 。

//输入：nums = [1,1,1,1,1]
//输出：[1,2,3,4,5]
//解释：动态和计算过程为 [1, 1+1, 1+1+1, 1+1+1+1, 1+1+1+1+1] 。

// runningSum[i] = sum(nums[0]…nums[i])

func main() {
	//nums := []int{1,2,3,4}
	nums := []int{1, 1, 1, 1, 1}

	fmt.Println(runningSum(nums))
}

func runningSum(nums []int) []int {
	newArr := make([]int, len(nums))

	sum := 0
	for i := range nums {
		sum = sum + nums[i]
		newArr[i] = sum
	}
	return newArr
}
