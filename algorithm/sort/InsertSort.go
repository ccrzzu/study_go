package sort

import "fmt"

/**
插入排序
时间复杂度 o(n^2)
空间复杂度 o(1)
原理是，从第二个数开始向右侧遍历，
每次均把该位置的元素移动至左侧，
放在放在一个正确的位置（比左侧大，比右侧小）。
*/

func InserSort(nums []int) {
	for i := 0; i < len(nums); i++ {
		insertVal := nums[i]
		insertIndex := i - 1
		for insertIndex >= 0 && insertVal < nums[insertIndex] {
			nums[insertIndex+1] = nums[insertIndex]
			insertIndex--
		}

		if insertIndex+1 != i {
			nums[insertIndex+1] = insertVal
		}
	}
	fmt.Println(nums)
}
