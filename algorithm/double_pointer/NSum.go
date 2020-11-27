package double_pointer

import "sort"

//两数相加等于目标数
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

//两数之和，输入的数组是有序的，下标从1开始
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

//返回值，所有满足的
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

func threeSum(nums []int) [][]int {
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

func threeSumNoRepeat(nums []int, start, target int) [][]int {
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

func fourSum(nums []int, target int) [][]int {
	sort.Ints(nums)
	res := [][]int{}
	// 穷举 fourSum 的第一个数
	for i := 0; i < len(nums); i++ {
		repeat := threeSumNoRepeat(nums, i+1, target-nums[i])
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
