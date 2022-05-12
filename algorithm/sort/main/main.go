package main

import (
	"fmt"
	"study_go/algorithm/sort"
)

func main() {
	nums := []int{4, 3, 2, 7, 8, 2, 3, 1}
	//cyclic_sort.CyclicSort(nums)
	//sort.BubbleSort(nums)
	//sort.SelectSort(nums)
	//sort.InserSort(nums)
	fmt.Println("----------------")
	sort.MergeSort(nums)
}
