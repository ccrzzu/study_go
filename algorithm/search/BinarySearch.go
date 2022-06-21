package search

import (
	"fmt"
	"sort"
)

// 给出⼀个有序数组 nums 和⼀个数 target ，要求在数组中找到第⼀个和这个元素相等的元素下
// 标，最后⼀个和这个元素相等的元素下标。

// 这⼀题是经典的⼆分搜索变种题。⼆分搜索有 4 ⼤基础变种题：
// 1. 查找第⼀个值等于给定值的元素
// 2. 查找最后⼀个值等于给定值的元素
// 3. 查找第⼀个⼤于等于给定值的元素
// 4. 查找最后⼀个⼩于等于给定值的元素

//普通二分查找
func binarySearch(nums []int, target int) int {
	left, right := 0, len(nums)-1
	for left <= right {
		mid := left + (right-left)/2
		if nums[mid] == target {
			return mid
		} else if nums[mid] > target {
			right = mid - 1
		} else if nums[mid] < target {
			left = mid + 1
		}
	}
	return -1
}

// 1. 查找第⼀个值等于给定值的元素
func searchFirstEqualElement(nums []int, target int) int {
	low, high := 0, len(nums)-1
	for low <= high {
		mid := low + (high-low)/2
		if nums[mid] > target {
			high = mid - 1
		} else if nums[mid] < target {
			low = mid + 1
		} else {
			if mid == 0 || nums[mid-1] < target {
				return mid
			}
			high = mid - 1
		}
	}
	return -1
}

//2. 查找最后⼀个值等于给定值的元素
func searchLastEqualElement(nums []int, target int) int {
	low, high := 0, len(nums)-1
	for low <= high {
		mid := low + (high-low)/2
		if nums[mid] > target {
			high = mid - 1
		} else if nums[mid] < target {
			low = mid + 1
		} else {
			if mid == len(nums)-1 || nums[mid+1] > target {
				return mid
			}
			low = mid + 1
		}
	}
	return -1
}



//二分查找所有符合的值
func binarySearchRange(nums []int, target int) []int {
	res := []int{}
	left, right := 0, len(nums)-1
	for left <= right {
		mid := left + (right-left)/2
		if nums[mid] == target {
			right = mid - 1
		} else if nums[mid] > target {
			right = mid - 1
		} else if nums[mid] < target {
			left = mid + 1
		}
	}
	if left == len(nums) || nums[left] != target {
		res = append(res, -1)
	} else {
		res = append(res, left)
	}
	left, right = 0, len(nums)-1
	for left <= right {
		mid := left + (right-left)/2
		if nums[mid] == target {
			left = mid + 1
		} else if nums[mid] > target {
			right = mid - 1
		} else if nums[mid] < target {
			left = mid + 1
		}
	}
	if right < 0 || nums[right] != target {
		res = append(res, -1)
	} else {
		res = append(res, right)
	}
	return res
}

// go内部的实现 面试肯定不能这么写 可以关注一下它的实现方式
func GoSearchInSort() {
	fmt.Println(sort.SearchInts([]int{1, 3, 6, 7, 8}, 7))
}
