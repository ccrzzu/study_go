package sliding_window

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
