package main

import (
	"fmt"
	"study_go/algorithm/array"
)

const cl = 100

var bl = 123

type Student struct {
	Name string
}

var list map[string]Student

func a() {
	list = make(map[string]Student)

	student := Student{"Aceld"}
	tmpStudent := student
	tmpStudent.Name = "LDB"
	list["student"] = tmpStudent

	cl := 1
	fmt.Println(cl)
}

func main() {
	// println(&bl,bl)
	// //println(&cl,cl)
	// a := "wo爱你"
	// fmt.Println(len(a))
	// fmt.Println(len([]byte(a)))
	// fmt.Println(len([]rune(a)))
	// fmt.Println(utf8.RuneCountInString(a))
	// fmt.Println(string([]rune(a)[2:3]))
	// // for _, item := range []byte("abc") {
	// // 	fmt.Println(item)
	// // }
	// // for _, item := range "abc" {
	// // 	fmt.Println(item - 'a')
	// // }
	// //nums := []int{1,1,1,2,2,3}
	// //array.RemoveDuplicatesWithTwoDuplicate(nums)
	// //digits := []int{1, 2, 0, 0}
	// //array.PlusOne2(digits)
	// //array.AddToArrayForm(digits, 34)
	// array.ThirdMax([]int{5, 2, 2})
	// b := []int{3, 4, 2, 1}
	// sort.Ints(b)
	// fmt.Println(b)
	// n := 23123
	// nums := []int{2, 4, 9}
	// fmt.Println(array.LargestNumberLessN(nums, n))

	arr := []int{3379, 6319, 8468, 10101, 11334, 12478, 13568, 14586, 15512, 16479, 24337, 30907, 36500,59301}
	res := array.CalculateDifferences(arr)
	fmt.Println(res)
}
