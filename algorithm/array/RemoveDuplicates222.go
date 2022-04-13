package array

import "fmt"

func removeDuplicates2(nums []int) int {
	if len(nums) <= 1 {
		return len(nums)
	}
	index := 1
	for i := 0; i < len(nums)-1; i++ {
		if nums[i+1] != nums[i] {
			nums[index] = nums[i+1]
			index++
		}
	}
	// result is nums[:index]
	return index
}

// 和上题一样，但是可以允许重复两次。
func RemoveDuplicatesWithTwoDuplicate(nums []int) int {
	if len(nums) <= 2 {
		return len(nums)
	}
	// index := 2
	// for i := 0; i < len(nums)-2; i++ {
	// 	if nums[i+2] != nums[i] {
	// 		nums[index] = nums[i+2]
	// 		index += 1
	// 	}
	// }
	index := 0
	for _, item := range nums {
		if index < 2 || item > nums[index-2] {
			nums[index] = item
			fmt.Println(index,nums)
			index += 1
		}
	}
	return index
}
