package stack

import "container/heap"

//优先级队列的应用

/**
 * 215,找到一个数组里的第K大值
 */
func findKthLargest(nums []int, k int) int {
	pq := make(PriorityQueue, len(nums))
	for i := 0; i < len(nums); i++ {
		pq[i] = &Item{
			value:    nums[i],
			priority: nums[i],
			index:    i,
		}
	}
	heap.Init(&pq)
	var i int
	for pq.Len() > 0 {
		i++
		v := heap.Pop(&pq).(*Item)
		if i == k {
			return v.value
		}
	}
	return -1
}

/**
 *88.合并2个有序链表
 */
func merge(nums1 []int, m int, nums2 []int, n int) {
	pq := make(PriorityQueue, m+n)
	i := 0
	for ; i < m; i++ {
		pq[i] = &Item{
			value:    nums1[i],
			priority: nums1[i],
			index:    i,
		}
	}
	for ; i < m+n; i++ {
		pq[i] = &Item{
			value:    nums2[i],
			priority: nums2[i],
			index:    i,
		}
	}
	heap.Init(&pq)
	var j int
	for pq.Len() > 0 {
		v := heap.Pop(&pq).(*Item)
		nums1[j] = v.value
		j++
	}
}

type Item struct {
	value    int // 优先级队列中的数据，可以是任意类型，这里使用int
	priority int // 优先级队列中节点的优先级
	index    int // index是该节点在堆中的位置
}

// 优先级队列需要实现heap的interface
type PriorityQueue []*Item

// 绑定Len方法
func (pq PriorityQueue) Len() int {
	return len(pq)
}

// 绑定Less方法，这里用的是小于号，生成的是小根堆
func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i].priority < pq[j].priority
}

// 绑定swap方法
func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index, pq[j].index = i, j
}

// 绑定put方法，将index置为-1是为了标识该数据已经出了优先级队列了
func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	*pq = old[0 : n-1]
	item.index = -1
	return item
}

// 绑定push方法
func (pq *PriorityQueue) Push(x interface{}) {
	n := len(*pq)
	item := x.(*Item)
	item.index = n
	*pq = append(*pq, item)
}

// 更新修改了优先级和值的item在优先级队列中的位置
func (pq *PriorityQueue) update(item *Item, value int, priority int) {
	item.value = value
	item.priority = priority
	heap.Fix(pq, item.index)
}
