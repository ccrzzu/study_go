package array

import "strconv"

/**
给定一个无重复元素的有序整数数组 nums 。

返回 恰好覆盖数组中所有数字 的 最小有序 区间范围列表。也就是说，nums 的每个元素都恰好被某个区间范围所覆盖，并且不存在属于某个范围但不属于 nums 的数字 x 。

列表中的每个区间范围 [a,b] 应该按如下格式输出：

“a->b” ，如果 a != b
“a” ，如果 a == b
*/
func summaryRanges(nums []int) []string {
	ans := []string{}
	for i, n := 0, len(nums); i < n; {
		left := i
		for i = i + 1; i < n && nums[i] == nums[i-1]+1; i++ {
		}
		s := strconv.Itoa(nums[left])
		//当left不是原地为动，有连续的数时
		if left != i-1 {
			s = s + "->" + strconv.Itoa(nums[i-1])
		}
		ans = append(ans, s)
	}
	return ans
}
