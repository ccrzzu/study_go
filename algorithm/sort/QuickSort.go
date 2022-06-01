package sort

import "fmt"

/**
快速排序
时间复杂度 o(nlogn)
空间复杂度 o(logn)
原理是，分治思想
1、其实就是把一个数组里的某位元素(本文是数组第一位)挪到正确的位置上，
   使得该元素左边全部小于它，右边元素全部大于它，这个过程也称为partition（分治）。
2、然后再分别对左右两部分数组分别再进行partition,如此进行递归下去…
3、直到左右两部分的子数组只含一个元素就结束递归。
*/
func QuickSort(nums []int) {
	QuicSortDG(nums, 0, len(nums)-1)
	fmt.Println(nums)
}

func QuicSortDG(nums []int, left, right int) {
	if left < right {
		partIndex := partition(nums, left, right)
		QuicSortDG(nums, left, partIndex-1)
		QuicSortDG(nums, partIndex+1, right)
	}
}

func partition(nums []int, left, right int) int {
	//这个基准值随意找，一般找第一个数
	pivotKeyVal := nums[left]
	for left < right {
		//先从right开始找，比基准值大的，就继续
		for left < right && nums[right] >= pivotKeyVal {
			right--
		}
		//right值小于基准值了，就把right的值放到left位置
		nums[left] = nums[right]

		//再从left开始找，比基准值小的，就继续
		for left < right && nums[left] <= pivotKeyVal {
			left++
		}
		//left值大于基准值了，就把left的值放到right位置
		nums[right] = nums[left]
	}
	// 当left和right相等时，相当于已经完成了一轮排序，把基准值pivotKeyVal赋值给中间
	nums[left] = pivotKeyVal
	return left
}
