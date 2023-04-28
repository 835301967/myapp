package easy_model

import (
	"fmt"
	"testing"
)

func TestTwoSum(t *testing.T) {
	sum := twoSum([]int{1, 2, 3, 4, 6, 7, 9}, 10)
	fmt.Println(sum)
}

func twoSum(nums []int, target int) []int {
	resp := make([]int, 0)
	numsMap := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		if v, ok := numsMap[target-nums[i]]; ok {
			resp = append(resp, v)
			resp = append(resp, i)
		}
		numsMap[nums[i]] = i
	}
	return resp
}
