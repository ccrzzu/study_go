package stack

import (
	"container/heap"
	"fmt"
)

//优先级队列的应用

/**
 * 215,找到一个数组里的第K大值
 */
func findKthLargest(nums []int, k int) int {
	pq := make(PriorityQueue, len(nums))
	for i := 0; i < len(nums); i++ {
		pq[i] = &Item{
			Value:    nums[i],
			Priority: nums[i],
			Index:    i,
		}
	}
	heap.Init(&pq)
	var i int
	for pq.Len() > 0 {
		i++
		v := heap.Pop(&pq).(*Item)
		if i == k {
			return v.Value
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
			Value:    nums1[i],
			Priority: nums1[i],
			Index:    i,
		}
	}
	for ; i < m+n; i++ {
		pq[i] = &Item{
			Value:    nums2[i],
			Priority: nums2[i],
			Index:    i,
		}
	}
	heap.Init(&pq)
	var j int
	for pq.Len() > 0 {
		v := heap.Pop(&pq).(*Item)
		nums1[j] = v.Value
		j++
	}
}

type Item struct {
	Value    int // 优先级队列中的数据，可以是任意类型，这里使用int
	Priority int // 优先级队列中节点的优先级
	Index    int // index是该节点在堆中的位置
}

// 优先级队列需要实现heap的interface
type PriorityQueue []*Item

// 绑定Len方法
func (pq PriorityQueue) Len() int {
	return len(pq)
}

// 绑定Less方法，这里用的是小于号，生成的是小根堆，大于号生成的是大根堆
func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i].Priority < pq[j].Priority
}

// 绑定swap方法
func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].Index, pq[j].Index = i, j
}

// 绑定put方法，将index置为-1是为了标识该数据已经出了优先级队列了
func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	*pq = old[0 : n-1]
	item.Index = -1
	return item
}

// 绑定push方法
func (pq *PriorityQueue) Push(x interface{}) {
	n := len(*pq)
	item := x.(*Item)
	item.Index = n
	*pq = append(*pq, item)
}

// 更新修改了优先级和值的item在优先级队列中的位置
func (pq *PriorityQueue) update(item *Item, value int, priority int) {
	item.Value = value
	item.Priority = priority
	heap.Fix(pq, item.Index)
}

//虾皮 新加坡

//  merge mulit sorted list into one sorted array.
func main() {
	l1 := &SortedList{Nums: []int{1, 4, 7}}
	l2 := &SortedList{Nums: []int{2, 5, 8}}
	sortedListArr := []*SortedList{l1, l2}
	fmt.Println(mergeMultiSortedList(sortedListArr))
}

type IntHeap []int // 定义一个类型

func (h IntHeap) Len() int { return len(h) } // 绑定len方法,返回长度

func (h IntHeap) Less(i, j int) bool { // 绑定less方法
	return h[i] < h[j] // 如果h[i]<h[j]生成的就是小根堆，如果h[i]>h[j]生成的就是大根堆
}

func (h IntHeap) Swap(i, j int) { // 绑定swap方法，交换两个元素位置
	h[i], h[j] = h[j], h[i]
}

func (h *IntHeap) Pop() interface{} { // 绑定pop方法，从最后拿出一个元素并返回
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func (h *IntHeap) Push(x interface{}) { // 绑定push方法，插入新元素
	*h = append(*h, x.(int))
}

type SortedList struct {
	Nums []int
}

func mergeMultiSortedList(l []*SortedList) []int {
	res := []int{}
	h := &IntHeap{}
	heap.Init(h)
	for _, item := range l {
		for _, num := range item.Nums {
			heap.Push(h, num)
		}
	}
	for len(*h) > 0 {
		n := heap.Pop(h).(int)
		res = append(res, n)
	}
	return res
}
