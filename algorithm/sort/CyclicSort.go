package sort

import "sort"

/**
 * 循环排序
 * 具体来说，我们遍历数组的每一位数字，如果当前数字不在正确的索引上，
 * 则将其与正确的索引交换，如下图所示。如果直接把每个数字放到正确的索引上，
 * 会产生平方级的时间复杂度，
 * 重点是：循环排序模式则可以提供线性的时间复杂度。
 */

//找到所有数组中消失的数字
func FindDisappearedNumbers(nums []int) []int {
	i, n := 0, len(nums)
	for i < n {
		j := nums[i] - 1
		if j < n && nums[i] != nums[j] {
			nums[i], nums[j] = nums[j], nums[i]
		} else {
			i++
		}
	}
	res := []int{}
	for i := 0; i < n; i++ {
		if nums[i] != i+1 {
			res = append(res, i+1)
		}
	}
	return res
}

/**
268
给定一个包含 [0, n] 中 n 个数的数组 nums （不一定有序，所以做题时需要先排序），
找出 [0, n] 这个范围内没有出现在数组中的那个数。
*/
func MissingNumber(nums []int) int {
	i, n := 0, len(nums)
	// 这一步是排序，
	for i < n {
		j := nums[i]
		if j < n && nums[i] != nums[j] {
			nums[i], nums[j] = nums[j], nums[i]
		} else {
			i++
		}
	}
	// 这一步是判断
	for i := 0; i < n; i++ {
		if nums[i] != i {
			return i
		}
	}
	return n
}

//解法二 利用位运算原理 异或 1^1=0 0^0=0 1^0=1 0^1=1
func MissingNumber2(nums []int) int {
	res, i := 0, 0
	for ; i < len(nums); i++ {
		res = res ^ i ^ nums[i]
	}
	return res ^ i
}

//给你一个未排序的整数数组，请你找出其中没有出现的最小的正整数。
//利用循环数组
func FirstMissingPositive(nums []int) int {
	i, n := 0, len(nums)
	for i < n {
		if nums[i] <= 0 {
			i++
			continue
		}
		j := nums[i] - 1
		if j < n && nums[i] != nums[j] {
			nums[i], nums[j] = nums[j], nums[i]
		} else {
			i++
		}
	}
	for i := 0; i < n; i++ {
		if nums[i] != i+1 {
			return i + 1
		}
	}
	return n + 1
}

//给你一个未排序的整数数组，请你找出其中没有出现的最小的正整数。
//利用map先装下原始数据
func FirstMissingPositive2(nums []int) int {
	m := make(map[int]bool)
	for i := 0; i < len(nums); i++ {
		m[nums[i]] = true
	}
	for i := 0; i < len(nums); i++ {
		if !m[i+1] {
			return i + 1
		}
	}
	return len(nums) + 1
}

//给定一个包含n + 1 个整数的数组 nums，其数字都在 1 到 n 之间（包括 1 和 n），可知至少存在一个重复的整数。假设只有一个重复的整数，找出这个重复的数。
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

//给定一个整数数组 a，其中1 ≤ a[i] ≤ n （n为数组长度）, 其中有些元素出现两次而其他元素出现一次。
//找到所有出现两次的元素。
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

//集合 S 包含从1到 n 的整数。不幸的是，因为数据错误，导致集合里面某一个元素复制了成了集合里面的另外一个元素的值，导致集合丢失了一个整数并且有一个元素重复。
//
//给定一个数组 nums 代表了集合 S 发生错误后的结果。你的任务是首先寻找到重复出现的整数，再找到丢失的整数，将它们以数组的形式返回。
func FindErrorNums(nums []int) []int {
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
			res = append(res, i+1)
			return res
		}
	}
	return res
}

//在一个长度为 n 的数组 nums 里的所有数字都在 0～n-1 的范围内。数组中某些数字是重复的，但不知道有几个数字重复了，也不知道每个数字重复了几次。请找出数组中任意一个重复的数字。
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

//给定一个非空整数数组，除了某个元素只出现一次以外，其余每个元素均出现两次。找出那个只出现了一次的元素。
func SingleNumber(nums []int) int {
	m := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		m[nums[i]]++
	}
	for k, v := range m {
		if v == 1 {
			return k
		}
	}
	return -1
}
