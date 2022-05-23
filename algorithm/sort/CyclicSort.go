package sort

import "fmt"

/**
 * 循环排序
 * 具体来说，我们遍历数组的每一位数字，如果当前数字不在正确的索引上，则将其与正确的索引交换。
 * 重点是：循环排序模式则可以提供线性的时间复杂度。
 */

//题设
//给你一个含 n 个整数的数组 nums ，其中 nums[i] 在区间 [1, n] 内,(或在[0,n-1]内)，
//那么对该数组进行排序
func CyclicSort(nums []int) {
	i, n := 0, len(nums)
	for i < n {
		fmt.Println("i:", i)
		j := nums[i] - 1
		fmt.Println("j:", j)
		if j < n && nums[i] != nums[j] {
			nums[i], nums[j] = nums[j], nums[i]
		} else {
			i++
		}
		fmt.Println(nums)
	}
	fmt.Println(nums)
}
