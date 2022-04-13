package array

/**
 * 给定一个大小为 n 的数组，找到其中的众数。
 * 众数是指在数组中出现次数大于 ⌊ n/2 ⌋ 的元素。你可以假设数组是非空的，并且给定的数组总是存在众数。
 *
解题思路 #
题目要求找出数组中出现次数大于 ⌊ n/2 ⌋ 次的数。要求空间复杂度为 O(1)。简单题。
这一题利用的算法是（波义尔摩尔大多数投票算法：Boyer-Moore Majority Vote Algorithm。）
https://www.zhihu.com/question/49973163/answer/235921864
*/

// 解法一 时间复杂度 O(n) 空间复杂度 O(1)
func MajorityElement(nums []int) int {
	majorityRes, count := nums[0], 0
	for i := 0; i < len(nums); i++ {
		if count == 0 {
			majorityRes, count = nums[i], 1
		} else {
			if majorityRes != nums[i] {
				count--
			} else {
				count++
			}
		}
	}
	return majorityRes
}

// 解法二 借助map，空间复杂度o(n)
func MajorityElement2(nums []int) int {
	m := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		m[nums[i]]++
		if m[nums[i]] > len(nums)/2{
			return nums[i]
		}
	}
	return 0
}

//给定一个大小为 n 的数组，找出其中所有出现超过 ⌊ n/3 ⌋ 次的元素。说明: 要求算法的时间复杂度为 O(n)，空间复杂度为 O(1)。
func MajorityElement229(nums []int) []int{
	// since we are checking if a num appears more than 1/3 of the time
	// it is only possible to have at most 2 nums (>1/3 + >1/3 = >2/3)
	count1, count2, candidate1, candidate2 := 0, 0, 0, 1
	// Select Candidates
	for _, num := range nums {
		if num == candidate1 {
			count1++
		} else if num == candidate2 {
			count2++
		} else if count1 <= 0 {
			// We have a bad first candidate, replace!
			candidate1, count1 = num, 1
		} else if count2 <= 0 {
			// We have a bad second candidate, replace!
			candidate2, count2 = num, 1
		} else {
			// Both candidates suck, boo!
			count1--
			count2--
		}
	}
	// Recount!
	count1, count2 = 0, 0
	for _, num := range nums {
		if num == candidate1 {
			count1++
		} else if num == candidate2 {
			count2++
		}
	}
	length := len(nums)
	if count1 > length/3 && count2 > length/3 {
		return []int{candidate1, candidate2}
	}
	if count1 > length/3 {
		return []int{candidate1}
	}
	if count2 > length/3 {
		return []int{candidate2}
	}
	return []int{}
}
