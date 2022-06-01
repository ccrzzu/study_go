package array

import (
	"container/heap"
	"study_go/algorithm/sort"
	"study_go/algorithm/stack"
)

/**
 * 215,找到一个数组里的第K大值
 */

//先排序的解法
func findKthLargest(nums []int, k int) int {
	sort.QuickSort(nums)
	return nums[len(nums)-k]
}

//大根堆解法，就是优先级队列解法
func findKthLargest2(nums []int, k int) int {
	pq := make(stack.PriorityQueue, len(nums))
	for i := 0; i < len(nums); i++ {
		pq[i] = &stack.Item{
			Value:    nums[i],
			Priority: nums[i],
			Index:    i,
		}
	}
	heap.Init(&pq)
	var i int
	for pq.Len() > 0 {
		i++
		v := heap.Pop(&pq).(*stack.Item)
		if i == k {
			return v.Value
		}
	}
	return -1
}
