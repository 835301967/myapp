package two_pointer_slice

import (
	"fmt"
	"testing"
)

/**
给定一个字符串 s ，请你找出其中不含有重复字符的 最长子串 的长度
题目：https://leetcode.cn/problems/longest-substring-without-repeating-characters/
*/

func TestLengthOfLongestSubString(t *testing.T) {
	substring1 := lengthOfLongestSubstring("abcabcbb")
	substring2 := lengthOfLongestSubstring("bbbbb")
	substring3 := lengthOfLongestSubstring("pwwkew")
	fmt.Println(substring1)
	fmt.Println(substring2)
	fmt.Println(substring3)
}

func lengthOfLongestSubstring(s string) int {
	max := 0
	dataMap := make(map[uint8]int)
	start := 0
	for end := 0; end < len(s); end++ {
		if v, ok := dataMap[s[end]]; ok {
			if v+1 > start {
				start = v + 1
			}
		}
		if end-start+1 > max {
			max = end - start + 1
		}

		dataMap[s[end]] = end
	}
	return max
}
