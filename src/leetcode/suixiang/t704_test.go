package suixiang

import (
	"fmt"
	"testing"
)

//输入: nums = [-1,0,3,5,9,12], target = 9  升序  、不重复的
//输出: 4
//解释: 9 出现在 nums 中并且下标为 4

func Test704(t *testing.T) {
	nums := []int{-1, 0, 3, 5, 9, 12, 13, 14, 16, 17, 19}
	fmt.Println(search(nums, 9))
}

func search(nums []int, target int) int {
	i := search2(nums, target)
	if i == -1 {
		return -1
	} else {
		return i - 1
	}

}
func search2(nums []int, target int) int {
	if len(nums) == 1 {
		if nums[0] == target {
			return 1
		} else {
			return -1
		}
	}
	if len(nums) <= 2 {
		if nums[0] == target {
			return 1
		} else if nums[1] == target {
			return 2
		} else {
			return -1
		}

	}
	mid := len(nums)/2 + 1

	if nums[mid-1] > target {
		m := search2(nums[:mid], target)
		if m == -1 {
			return -1
		}
		return m
	} else if nums[mid-1] < target {
		m := search2(nums[mid:], target)
		if m == -1 {
			return -1
		}
		return m + mid
	} else {
		return mid
	}
	return -1
}
