package sort

import "fmt"

/**
归并排序
时间复杂度 o(nlogn)
空间复杂度 o(n)
原理也是分治思想
与快速排序一样，归并排序采用的也是分治的策略，把原本的问题先分解成一些小问题进行求解，再把这些小问题各自的答案修整到一起得到原本问题的答案，从而达到分而治之的目的。
1、分割：归并排序算法会把要排序的序列分成长度相当的两个子序列，
	当分无可分每个子序列中只有一个数据的时候，就对子序列进行归并。
2、归并：指的是把两个排序好的子序列合并成一个有序序列。
	该操作会一直重复执行，直到所有子序列归并为一个整体为止。
*/

func MergeSort(nums []int) {
	mergeSort(nums, 0, len(nums)-1)
	fmt.Println(nums)
}

func mergeSort(nums []int, left, right int) {
	if left >= right {
		//此时不可再分割，因为只有一个元素了
		return
	}
	mid := (left + right) / 2
	mergeSort(nums, left, mid)
	mergeSort(nums, mid+1, right)
	if nums[mid] > nums[mid+1] {
		merge(nums, left, mid, right)
	}
}

//将nums[left...mid]和nums[mid+1...right]的两个区间进行归并
func merge(nums []int, left, mid, right int) {
	//先将left到right区间的元素copy到一个临时数组
	tempArr := make([]int, right+1)
	for index := left; index <= right; index++ {
		tempArr[index] = nums[index]
	}

	//初始化要归并的两个区间的起始坐标
	i := left
	j := mid + 1

	//遍历并逐个确定数组里left到right内的值的正确顺序
	for k := left; k <= right; k++ {
		//首先要注意i的边界[left...mid],j的边界[mid+1...right]
		if i > mid {
			nums[k] = tempArr[j]
			j++
			continue
		}
		if j > right {
			nums[k] = tempArr[i]
			i++
			continue
		}

		//不是上述两者情况，则开始比较两个区间数组的元素
		if tempArr[i] < tempArr[j] {
			nums[k] = tempArr[i]
			i++
		} else {
			nums[k] = tempArr[j]
			j++
		}
	}
}
