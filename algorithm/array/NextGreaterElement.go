package array

/**
nums1 中数字 x 的 下一个更大元素 是指 x 在 nums2 中对应位置 右侧 的 第一个 比 x 大的元素。

给你两个 没有重复元素 的数组 nums1 和 nums2 ，下标从 0 开始计数，其中nums1 是 nums2 的子集。

对于每个 0 <= i < nums1.length ，找出满足 nums1[i] == nums2[j] 的下标 j ，并且在 nums2 确定 nums2[j] 的 下一个更大元素 。如果不存在下一个更大元素，那么本次查询的答案是 -1 。

返回一个长度为 nums1.length 的数组 ans 作为答案，满足 ans[i] 是如上所述的 下一个更大元素 。
*/
func NextGreaterElement(nums1, nums2 []int) []int {
	ans := []int{}
	if len(nums1) == 0 || len(nums2) == 0 {
		return ans
	}
	nums2Map := make(map[int]int, 0)
	for i, item := range nums2 {
		nums2Map[item] = i
	}
	for i := 0; i < len(nums1); i++ {
		flag := false
		for j := nums2Map[nums1[i]]; j < len(nums2); j++ {
			if nums2[j] > nums1[i] {
				ans = append(ans, nums2[j])
				flag = true
				break
			}
		}
		if !flag {
			ans = append(ans, -1)
		}
	}
	return ans
}
