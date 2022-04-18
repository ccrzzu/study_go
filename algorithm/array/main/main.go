package main

import (
	"fmt"
	"sort"
	"study_go/algorithm/array"
)

func main() {
	// for _, item := range []byte("abc") {
	// 	fmt.Println(item)
	// }
	// for _, item := range "abc" {
	// 	fmt.Println(item - 'a')
	// }
	//nums := []int{1,1,1,2,2,3}
	//array.RemoveDuplicatesWithTwoDuplicate(nums)
	//digits := []int{1, 2, 0, 0}
	//array.PlusOne2(digits)
	//array.AddToArrayForm(digits, 34)
	array.ThirdMax([]int{5,2,2})
	a := []int{3,4,2,1}
	sort.Ints(a)
	fmt.Println(a)
}
