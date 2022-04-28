package array

//数组中是否包含重复元素，且元素下标差小于等于k。
func containsNearbyDuplicate(nums []int, k int) bool {
	record := make(map[int]int, len(nums))
	for i, item := range nums {
		if j, ok := record[item]; ok {
			if i-j <= k {
				return true
			}
		}
		record[item] = i
	}
	return false
}

//给你一个整数数组 nums 和两个整数 k 和 t 。请你判断是否存在 两个不同下标 i 和 j，
//使得 abs(nums[i] - nums[j]) <= t ，同时又满足 abs(i - j) <= k 。
// 解法一  滑动窗口 + 剪枝  
// 还有一种解法是桶排序 ：https://books.halfrost.com/leetcode/ChapterFour/0200~0299/0220.Contains-Duplicate-III/
func containsNearbyAlmostDuplicate(nums []int, k int, t int) bool {
	if len(nums) <= 1 {
		return false
	}
	if k <= 0 {
		return false
	}
	n := len(nums)
	for i := 0; i < n; i++ {
		count := 0
		for j := i + 1; j < n && count < k; j++ {
			if abs(nums[i]-nums[j]) <= t {
				return true
			}
			count++
		}
	}
	return false
}
