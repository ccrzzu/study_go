package array

import (
	"math"
	"strconv"
	"strings"
)

// LIS（Longest Increasing Subsequence）
// 最长递增子序列一个数的序列bi，当b1 < b2 < … < bS的时候，我们称这个序列是上升的

//给定一个未排序的整数数组，找到最长递增子序列的个数。
func FindNumberOfLIS(nums []int) int {
	dp := make([]int, len(nums), len(nums))
	count := make([]int, len(nums), len(nums))
	res := 0
	maxLen := 0
	for i := 0; i < len(nums); i++ {
		dp[i] = 1
		count[i] = 1
		for j := 0; j < i; j++ {
			if nums[j] < nums[i] {
				//dp[i] = int(math.Max(float64(dp[i]), float64(dp[j]+1)))
				if dp[j]+1 == dp[i] {
					count[i] += count[j]
				}
				if dp[j]+1 > dp[i] {
					dp[i] = dp[j] + 1
					count[i] = count[j]
				}
			}
		}
		if dp[i] == maxLen {
			res += count[i]
		}
		if dp[i] > maxLen {
			maxLen = dp[i]
			res = count[i]
		}
	}
	return res
}

//给定一个无序的整数数组，找到其中最长递增子序列的长度。
func LengthOfLIS(nums []int) int {
	dp := make([]int, len(nums), len(nums))
	maxLen := 0
	for i := 0; i < len(nums); i++ {
		dp[i] = 1
		for j := 0; j < i; j++ {
			if nums[j] < nums[i] {
				dp[i] = int(math.Max(float64(dp[i]), float64(dp[j]+1)))
			}
		}
		if dp[i] > maxLen {
			maxLen = dp[i]
		}
	}
	return maxLen
}

//给定一个整型数组, 你的任务是找到所有该数组的递增子序列，递增子序列的长度至少是2。
func FindSubsequences(nums []int) [][]int {
	resMap := make(map[int][][]int)
	for i := 0; i < len(nums); i++ {
		tmp := [][]int{{nums[i]}}
		for k, v := range resMap {
			if nums[i] >= k {
				for _, item := range v {
					tmpItem := make([]int, 0)
					tmpItem = append(tmpItem, item...)
					tmpItem = append(tmpItem, nums[i])
					tmp = append(tmp, tmpItem)
				}
			}
		}
		resMap[nums[i]] = tmp
	}

	/*for _, v := range resMap {
		for i, item := range v {
			flag := true
			if i == 0 {
				continue
			}
			for _, item2 := range result {
				if testEq(item, item2) {
					flag = false
					break
				}
			}
			if flag {
				result = append(result, item)
			}
		}
	}*/
	resultMap := make(map[string]int)
	for _, v := range resMap {
		for i, item := range v {
			if i == 0 {
				continue
			}
			resultMap[SplitToString(item, ",")] = 0
		}
	}
	result := make([][]int, 0)
	for k := range resultMap {
		result = append(result, StringToIntSlice(k))
	}
	return result
}

func testEq(a, b []int) bool {
	// If one is nil, the other must also be nil.
	if (a == nil) != (b == nil) {
		return false
	}

	if len(a) != len(b) {
		return false
	}

	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}

	return true
}

func SplitToString(a []int, sep string) string {
	if len(a) == 0 {
		return ""
	}

	b := make([]string, len(a))
	for i, v := range a {
		b[i] = strconv.Itoa(v)
	}
	return strings.Join(b, sep)
}

func StringToIntSlice(r string) []int {
	res := make([]int, 0)
	splits := strings.Split(r, ",")
	for _, item := range splits {
		atoi, _ := strconv.Atoi(item)
		res = append(res, atoi)
	}
	return res
}

//LHS 最长和谐子序列
//和谐数组是指一个数组里元素的最大值和最小值之间的差别正好是1。
//现在，给定一个整数数组，你需要在所有可能的子序列中找到最长的和谐子序列的长度。
func FindLHS(nums []int) int {
	if len(nums) < 2 {
		return 0
	}
	var res int
	dpMap := make(map[int]int)
	for _, item := range nums {
		dpMap[item]++
	}
	for k, v := range dpMap {
		if dpMap[k+1] != 0 && dpMap[k+1]+v > res {
			res = dpMap[k+1] + v
		}
	}
	return res
}
