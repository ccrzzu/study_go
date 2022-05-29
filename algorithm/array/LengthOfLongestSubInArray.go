package array

import "math"

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
