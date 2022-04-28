package array

import (
	"math"
	"sort"
)

//1
//两数之和，相加等于目标数，输入的数组无序，下标从0开始
func TwoSum(nums []int, target int) []int {
	nMap := make(map[int]int)
	for i, item := range nums {
		if p, ok := nMap[target-item]; ok {
			return []int{i, p}
		}
		nMap[item] = i
	}
	return nil
}

// 解法二 两数之和，输入的数组是有序的，下标从1开始
func twoSum1(numbers []int, target int) []int {
	left, right := 0, len(numbers)-1
	for left < right {
		sum := numbers[left] + numbers[right]
		if sum == target {
			return []int{left + 1, right + 1}
		} else if sum < target {
			left++
		} else if sum > target {
			right--
		}
	}
	return []int{-1, -1}
}

// 两数之和，返回值是所有满足的值
func twoSum2(numbers []int, target int) [][]int {
	left, right := 0, len(numbers)-1
	res := [][]int{}
	for left < right {
		sum := numbers[left] + numbers[right]
		leftVal, rightVal := numbers[left], numbers[right]
		if sum == target {
			res = append(res, []int{numbers[left], numbers[right]})
			for left < right && numbers[left] == leftVal {
				left++
			}
			for left < right && numbers[right] == rightVal {
				right++
			}
		} else if sum < target {
			for left < right && numbers[left] == leftVal {
				left++
			}
		} else if sum > target {
			for left < right && numbers[right] == rightVal {
				right++
			}
		}
	}
	return res
}

// 两数之和，返回值是所有满足的值,且不重复
func twoSumNoRepeat(numbers []int, start, target int) [][]int {
	left, right := start, len(numbers)-1
	res := [][]int{}
	for left < right {
		sum := numbers[left] + numbers[right]
		leftVal, rightVal := numbers[left], numbers[right]
		if sum == target {
			res = append(res, []int{numbers[left], numbers[right]})
			for left < right && numbers[left] == leftVal {
				left++
			}
			for left < right && numbers[right] == rightVal {
				right--
			}
		} else if sum < target {
			for left < right && numbers[left] == leftVal {
				left++
			}
		} else if sum > target {
			for left < right && numbers[right] == rightVal {
				right--
			}
		}
	}
	return res
}

/**
 15
 给你一个包含 n 个整数的数组 nums，
 判断 nums 中是否存在三个元素 a，b，c ，
 使得 a + b + c = 0 ？请你找出所有和为 0 且不重复的三元组。
*/
func ThreeSum(nums []int) [][]int {
	sort.Ints(nums)
	res := [][]int{}
	// 穷举 threeSum 的第一个数
	for i := 0; i < len(nums); i++ {
		repeat := twoSumNoRepeat(nums, i+1, 0-nums[i])
		for _, item := range repeat {
			item = append(item, nums[i])
			res = append(res, item)
		}
		for i < len(nums)-1 && nums[i] == nums[i+1] {
			i++
		}
	}
	return res
}

// 15 三数之和，返回值没有重复的
func ThreeSumNoRepeat(nums []int, start, target int) [][]int {
	res := [][]int{}
	// 穷举 threeSum 的第一个数
	for i := start; i < len(nums); i++ {
		repeat := twoSumNoRepeat(nums, i+1, target-nums[i])
		for _, item := range repeat {
			item = append(item, nums[i])
			res = append(res, item)
		}
		for i < len(nums)-1 && nums[i] == nums[i+1] {
			i++
		}
	}
	return res
}

//18 给定一个数组，要求在这个数组中找出 4 个数之和为 0 的所有组合。
func FourSum(nums []int, target int) [][]int {
	sort.Ints(nums)
	res := [][]int{}
	// 穷举 fourSum 的第一个数
	for i := 0; i < len(nums); i++ {
		repeat := ThreeSumNoRepeat(nums, i+1, target-nums[i])
		for _, item := range repeat {
			item = append(item, nums[i])
			res = append(res, item)
		}
		for i < len(nums)-1 && nums[i] == nums[i+1] {
			i++
		}
	}
	return res
}

/**
16
给你一个长度为 n 的整数数组 nums 和 一个目标值 target。请你从 nums 中选出三个整数，使它们的和与 target 最接近。
返回这三个数的和。
*/
func ThreeSumClosest(nums []int, target int) int {
	n, res, diff := len(nums), 0, math.MaxInt32
	if n > 2 {
		sort.Ints(nums)
		for i := 0; i < n-2; i++ {
			//增加这段是为了去除重复计算项
			if i > 0 && nums[i] == nums[i-1] {
				continue
			}
			for j, k := i+1, n-1; j < k; {
				sum := nums[i] + nums[j] + nums[k]
				if abs(sum-target) < diff {
					res, diff = sum, abs(sum-target)
				}
				if sum == target {
					return res
				} else if sum > target {
					k--
				} else {
					j++
				}
			}
		}
	}
	return res
}

func abs(num int) int {
	if num >= 0 {
		return num
	}
	return -num
}
