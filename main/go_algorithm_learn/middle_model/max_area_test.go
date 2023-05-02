package middle_model

import (
	"fmt"
	"testing"
)

// 暴力解法
func maxArea(input []int) int {
	area := 0
	for i := 0; i < len(input)-1; i++ {
		for j := i + 1; j < len(input); j++ {
			height := input[j]
			if input[i] < input[j] {
				height = input[i]
			}
			tempArea := height * (j - i)
			if tempArea > area {
				area = tempArea
			}
		}
	}
	return area
}

// 双指针法解法
func maxAreaOpt(height []int) int {
	left := 0
	right := len(height) - 1
	area := 0
	for left < right {
		low := height[right]
		if height[left] < height[right] {
			low = height[left]
		}
		tempArea := low * (right - left)
		if tempArea > area {
			area = tempArea
		}
		if height[left] < height[right] {
			left++
		} else {
			right--
		}

	}
	return area
}

func TestGetMaxArea(t *testing.T) {
	exapmle1 := []int{1, 8, 6, 2, 5, 4, 8, 3, 7}
	area := maxAreaOpt(exapmle1)
	fmt.Println(area)
	exapmle2 := []int{1, 1}
	area2 := maxAreaOpt(exapmle2)
	fmt.Println(area2)
}
