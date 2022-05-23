package sort

/**
448
给你一个含 n 个整数的数组 nums ，其中 nums[i] 在区间 [1, n] 内。
请你找出所有在 [1, n] 范围内但没有出现在 nums 中的数字，并以数组的形式返回结果。
题解：读完题后 确定 首先nums是无序的，然后nums里有的元素一定是出现了2次，才会导致有数据丢失
*/
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
	// 这一步是循环排序，
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

//41
//给你一个未排序的整数数组，范围是[1,n], 请你找出其中没有出现的最小的正整数。
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

//41
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

//注意，这道题是FindDisappearedNumbers和FindDuplicate两道题的综合题
//集合 S 包含从1到 n 的整数。不幸的是，因为数据错误，
//导致集合里面某一个元素复制了成了集合里面的另外一个元素的值，导致集合丢失了一个整数并且有一个元素重复。
//
//给定一个数组 nums 代表了集合 S 发生错误后的结果。
//你的任务是首先寻找到重复出现的整数，再找到丢失的整数，将它们以数组的形式返回。
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
