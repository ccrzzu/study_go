package sort

/**
287
给定一个包含n + 1 个整数的数组 nums，其数字都在 1 到 n 之间（包括 1 和 n），
可知至少存在一个重复的整数。假设只有一个重复的整数，找出这个重复的数。
*/

//解法1 循环排序
func FindDuplicate(nums []int) int {
	i, n := 0, len(nums)
	for i < n {
		j := nums[i] - 1
		if j < n && nums[i] != nums[j] {
			nums[i], nums[j] = nums[j], nums[i]
		} else {
			i++
		}
	}
	for i := 0; i < n; i++ {
		if nums[i] != i+1 {
			return nums[i]
		}
	}
	return n + 1
}

/**
上题的变形题 ： 找到所有出现两次的元素。
给定一个整数数组 a，其中1 ≤ a[i] ≤ n （n为数组长度）, 其中有些元素出现两次而其他元素出现一次。
*/
func FindDuplicates(nums []int) []int {
	i, n := 0, len(nums)
	for i < n {
		j := nums[i] - 1
		if j < n && nums[i] != nums[j] {
			nums[i], nums[j] = nums[j], nums[i]
		} else {
			i++
		}
	}
	var res []int
	for i := 0; i < n; i++ {
		if nums[i] != i+1 {
			res = append(res, nums[i])
		}
	}
	return res
}

//在一个长度为 n 的数组 nums 里的所有数字都在 0～n-1 的范围内。
//数组中某些数字是重复的，但不知道有几个数字重复了，也不知道每个数字重复了几次。
//请找出数组中任意一个重复的数字。
func FindRepeatNumber(nums []int) int {
	i, n := 0, len(nums)
	for i < n {
		j := nums[i]
		if j < n && nums[i] != nums[j] {
			nums[i], nums[j] = nums[j], nums[i]
		} else {
			i++
		}
	}
	for i := 0; i < n; i++ {
		if nums[i] != i {
			return nums[i]
		}
	}
	return n
}

func FindRepeatNumber2(nums []int) int {
	m := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		m[nums[i]]++
	}
	for k, v := range m {
		if v >= 2 {
			return k
		}
	}
	return len(nums)
}
