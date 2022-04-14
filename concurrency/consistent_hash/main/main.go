package main

import (
	"fmt"
	"sort"
)

func main() {
	a := []int{1, 2, 3, 4, 5}
	b := sort.Search(len(a), func(i int) bool { return a[i] >= 30 })
	fmt.Println(b) //5，查找不到，返回a slice的长度５，而不是-1
	c := sort.Search(len(a), func(i int) bool { return a[i] <= 3 })
	fmt.Println(c) //0，利用二分法进行查找，返回符合条件的最左边数值的index，即为０
	d := sort.Search(len(a), func(i int) bool { return a[i] == 3 })
	fmt.Println(d)
	e := sort.Search(len(a), func(i int) bool { return a[i] >= 3 })
	fmt.Println(e) //0，利用二分法进行查找，返回符合条件的最右边数值的index，即为3
}
