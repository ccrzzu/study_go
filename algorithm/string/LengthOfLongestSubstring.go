package string

import (
	"math"
)

/**
 3
//最长无重复字符的子串的长度
//给定一个字符串，请你找出其中不含有重复字符的 最长子串 的长度。
//例如：s = "abcabcbb" 因为无重复字符的最长子串是 "abc"，所以其长度为 3。
*/
// 滑动窗口 map self
func LengthOfLongestSubstring1(s string) int {
	left, right, max := 0, 0, 0
	subMap := make(map[byte]int)
	for left <= right && right < len(s) {
		rightChar := s[right]
		subMap[rightChar]++
		right++
		for subMap[rightChar] > 1 {
			leftChar := s[left]
			left++
			subMap[leftChar]--
		}
		if right-left > max {
			max = right - left
		}
	}
	return max
}

// 滑动窗口 map 别人的 好理解
func LengthOfLongestSubstring2(s string) int {
	window := map[byte]int{}
	left, right, res := 0, 0, 0
	for right < len(s) {
		rightChar := s[right]
		right++
		window[rightChar]++
		for window[rightChar] > 1 {
			leftChar := s[left]
			left++
			window[leftChar]--
		}
		if right-left > res {
			res = right - left
		}
	}
	return res
}

// 滑动窗口 ASCII码 array
func LengthOfLongestSubstring3(s string) int {
	if len(s) == 0 {
		return 0
	}
	var freq [127]int
	result, left, right := 0, 0, 0
	for left < len(s) {
		if right < len(s) && freq[s[right]] == 0 {
			freq[s[right]]++
			right++
		} else {
			freq[s[left]]--
			left++
		}
		if right-left > result {
			result = right - left
		}
	}
	return result
}

// 位图
func LengthOfLongestSubstring4(s string) int {
	if len(s) == 0 {
		return 0
	}
	var bitSet [128]bool
	left, right, res := 0, 0, 0
	for left < len(s) {
		if !bitSet[s[right]] {
			bitSet[s[right]] = true
			right++
		} else {
			bitSet[s[left]] = false
			left++
		}
		if right-left > res {
			res = right - left
		}
		if left+res >= len(s) || right >= len(s) {
			break
		}
	}
	return res
}

// 滑动窗口 哈希桶
func LengthOfLongestSubstring5(s string) int {
	left, right, res := 0, 0, 0
	indexes := make(map[byte]int, len(s))
	for right < len(s) {
		if val, ok := indexes[s[right]]; ok && val >= left {
			left = val + 1
		}
		indexes[s[right]] = right
		right++
		if right-left > res {
			res = right - left
		}
	}
	return res
}
