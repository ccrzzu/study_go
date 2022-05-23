package sort

import "fmt"

/**
选择排序
时间复杂度 o(n^2)
空间复杂度 o(1)
原理是，对给定的数组进行多次遍历，每次均找出最大的一个值的索引。
*/
func SelectSort(nums []int) {
	n := len(nums)
	for i := 0; i < n; i++ {
		maxIndex := 0
		for j := 1; j < n-i; j++ {
			if nums[j] > nums[maxIndex] {
				maxIndex = j
			}
		}
		nums[n-i-1], nums[maxIndex] = nums[maxIndex], nums[n-i-1]
	}
	fmt.Println(nums)
}
