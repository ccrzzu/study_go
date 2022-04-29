package double_pointer

import "math"

//一个字符串包含另一个字符串所有字母的最小子串
func minWindow(s string, t string) string {
	window := map[byte]int{}
	need := map[byte]int{}
	for i := range t {
		need[t[i]]++
	}
	left, right, valid := 0, 0, 0
	start, length := 0, math.MaxInt32
	for right < len(s) {
		rightChar := s[right]
		right++
		if _, ok := need[rightChar]; ok {
			window[rightChar]++
			if need[rightChar] == window[rightChar] {
				valid++
			}
		}

		for valid == len(need) {
			if right-left < length {
				start = left
				length = right - left
			}
			leftChar := s[left]
			left++
			if _, ok := need[leftChar]; ok {
				if need[leftChar] == window[leftChar] {
					valid--
				}
				window[leftChar]--
			}
		}
	}
	if length != math.MaxInt32 {
		return s[start : start+length]
	}
	return ""
}

//判断 s2 中是否包含s1的排列
func checkInclusion(s1 string, s2 string) bool {
	window := map[byte]int{}
	need := map[byte]int{}
	for i, _ := range s1 {
		need[s1[i]]++
	}
	left, right, valid := 0, 0, 0
	for right < len(s2) {
		rightChar := s2[right]
		right++
		if _, ok := need[rightChar]; ok {
			window[rightChar]++
			if need[rightChar] == window[rightChar] {
				valid++
			}
		}

		for right-left == len(s1) {
			if valid == len(need) {
				return true
			}
			leftChar := s2[left]
			left++
			if _, ok := need[leftChar]; ok {
				if need[leftChar] == window[leftChar] {
					valid--
				}
				window[leftChar]--
			}
		}
	}
	return false
}

//找出所有字母异位词
func findAnagrams(s string, p string) []int {
	window := map[byte]int{}
	need := map[byte]int{}
	for i, _ := range p {
		need[p[i]]++
	}
	left, right, valid := 0, 0, 0
	res := []int{}
	for right < len(s) {
		rightChar := s[right]
		right++
		if _, ok := need[rightChar]; ok {
			window[rightChar]++
			if need[rightChar] == window[rightChar] {
				valid++
			}
		}

		for right-left == len(p) {
			if valid == len(need) {
				res = append(res, left)
			}
			leftChar := s[left]
			left++
			if _, ok := need[leftChar]; ok {
				if need[leftChar] == window[leftChar] {
					valid--
				}
				window[leftChar]--
			}
		}
	}
	return res
}

//最长无重复字符的子串
//给定一个字符串，请你找出其中不含有重复字符的 最长子串 的长度。
//例如：s = "abcabcbb" 因为无重复字符的最长子串是 "abc"，所以其长度为 3。
func lengthOfLongestSubstring(s string) int {
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
		res = int(math.Max(float64(res), float64(right-left)))
	}
	return res
}

//数组的最长无重复子串的长度
func lengthOfLongestSubInArray(arr []int) int {
	window := map[int]int{}
	left, right, res := 0, 0, 0
	for right < len(arr) {
		rightChar := arr[right]
		right++
		window[rightChar]++
		for window[rightChar] > 1 {
			leftChar := arr[left]
			left++
			window[leftChar]--
		}
		res = int(math.Max(float64(res), float64(right-left)))
	}
	return res
}

//滑动窗口内的最大值 暴力解法
func maxSlidingWindowByBaoLi(nums []int, k int) []int {
	len := len(nums)
	result := []int{}
	for i := 0; i < len-k+1; i++ {
		max := math.MinInt64
		for j := i; j < i+k; j++ {
			if nums[j] > max {
				max = nums[j]
			}
		}
		result = append(result, max)
	}
	return result
}

//滑动窗口内的最大值 双端队列解法
func maxSlidingWindow(nums []int, k int) []int {
	if len(nums) == 0 {
		return []int{}
	}
	queue := []int{}
	result := []int{}
	for i := range nums {
		for i > 0 && len(queue) > 0 && nums[i] > queue[len(queue)-1] {
			queue = queue[:len(queue)-1]
		}
		queue = append(queue, nums[i])
		if i >= k && nums[i-k] == queue[0] {
			queue = queue[1:]
		}
		if i >= k-1 {
			result = append(result, queue[0])
		}
	}
	return result
}
