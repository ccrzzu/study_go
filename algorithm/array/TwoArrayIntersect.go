package array

import "sort"

//两个数组的交集，结果中是需要包含重复次数的
// 给你两个整数数组 nums1 和 nums2 ，
//请你以数组形式返回两数组的交集。
//返回结果中每个元素出现的次数，应与元素在两个数组中都出现的次数一致
//（如果出现次数不一致，则考虑取较小值）。可以不考虑输出结果的顺序。

func Intersect(nums1 []int, nums2 []int) []int {
	mp0 := make(map[int]int)
	for _, item := range nums1 {
		mp0[item] += 1
	}
	k := 0
	for _, item := range nums2 {
		if mp0[item] > 0 {
			mp0[item] -= 1
			nums2[k] = item
			k++
		}
	}
	return nums2[0:k]
}

func Intersect2(nums1 []int, nums2 []int) []int {
	var res []int
	mp1 := make(map[int]int, 0)
	for _, item := range nums1 {
		mp1[item]++
	}
	for _, item := range nums2 {
		if mp1[item] > 0 {
			mp1[item]--
			res = append(res, item)
		}
	}
	return res
}

// 双指针策略对排好序的数组类型问题是一大招
// 让两个数组排序后，然后用双指针策略，
// 解答中我们并没有创建空白数组，因为遍历后的数组其实就没用了
// 我们可以将相等的元素放入用过的数组中，就为我们节省下了空间
func intersectAfterSortByTowPoint(nums1 []int, nums2 []int) []int {
	i, j, k := 0, 0, 0
	sort.Ints(nums1)
	sort.Ints(nums2)
	for i < len(nums1) && j < len(nums2) {
		if nums1[i] == nums2[j] {
			nums1[k] = nums2[j]
			i++
			j++
			k++
		} else if nums1[i] < nums2[j] {
			i++
		} else if nums2[j] < nums1[i] {
			j++
		}
	}
	return nums1[:k]
}

//让两个数组排序后，然后用双指针策略
// 解答中创建了空白数组
func intersectAfterSortByTowPointToNewMem(nums1 []int, nums2 []int) []int {
	i, j := 0, 0
	sort.Ints(nums1)
	sort.Ints(nums2)
	var res []int
	for i < len(nums1) && j < len(nums2) {
		if nums1[i] == nums2[j] {
			res = append(res, nums1[i])
			i++
			j++
		} else if nums1[i] < nums2[j] {
			i++
		} else if nums2[j] < nums1[i] {
			j++
		}
	}
	return res
}

//两个数组的交集，
//结果中重复次数只返回一次
func intersectUnique(nums1 []int, nums2 []int) []int {
	var res []int
	mp1 := make(map[int]int, 0)
	for _, item := range nums1 {
		mp1[item] = 0
	}
	for _, item := range nums2 {
		if _, ok := mp1[item]; ok {
			res = append(res, item)
			delete(mp1, item)
		}
	}
	return res
}
