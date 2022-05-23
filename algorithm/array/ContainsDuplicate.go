package array

//数组中是否包含重复元素，且元素下标差小于等于k。
func ContainsNearbyDuplicate(nums []int, k int) bool {
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

//这道题可以维护一个只有 K 个元素的 map，每次只需要判断这个 map 里面是否存在这个元素即可。
//如果存在就代表重复数字的下标差值在 K 以内。
//map 的长度如果超过了 K 以后就删除掉 i-k 的那个元素，这样一直维护 map 里面只有 K 个元素。
func containsNearbyDuplicate2(nums []int, k int) bool {
	if len(nums) <= 1 {
		return false
	}
	if k <= 0 {
		return false
	}
	record := make(map[int]bool, len(nums))
	for i, item := range nums {
		if _, ok := record[item]; ok {
			return true
		}
		record[item] = true
		if len(record) == k+1 {
			delete(record, nums[i-k])
		}
	}
	return false
}

//给你一个整数数组 nums 和两个整数 k 和 t 。请你判断是否存在 两个不同下标 i 和 j，
//使得 abs(nums[i] - nums[j]) <= t ，同时又满足 abs(i - j) <= k 。
// 解法一  滑动窗口 + 剪枝
// 还有一种解法是桶排序 ：https://books.halfrost.com/leetcode/ChapterFour/0200~0299/0220.Contains-Duplicate-III/
func ContainsNearbyAlmostDuplicate(nums []int, k int, t int) bool {
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

// 解法一 桶排序
func containsNearbyAlmostDuplicate2(nums []int, k int, t int) bool {
	if k <= 0 || t < 0 || len(nums) < 2 {
		return false
	}
	buckets := map[int]int{}
	for i := 0; i < len(nums); i++ {
		// Get the ID of the bucket from element value nums[i] and bucket width t + 1
		key := nums[i] / (t + 1)
		// -7/9 = 0, but need -7/9 = -1
		if nums[i] < 0 {
			key--
		}
		if _, ok := buckets[key]; ok {
			return true
		}
		// check the lower bucket, and have to check the value too
		if v, ok := buckets[key-1]; ok && nums[i]-v <= t {
			return true
		}
		// check the upper bucket, and have to check the value too
		if v, ok := buckets[key+1]; ok && v-nums[i] <= t {
			return true
		}
		// maintain k size of window
		if len(buckets) >= k {
			delete(buckets, nums[i-k]/(t+1))
		}
		buckets[key] = nums[i]
	}
	return false
}

// 解法二 滑动窗口 + 剪枝
func containsNearbyAlmostDuplicate3(nums []int, k int, t int) bool {
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
