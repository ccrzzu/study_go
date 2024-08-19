package main

import (
	"fmt"
	"strings"
)

func DeleteSlice(s []int, elem int) []int {
	for i := 0; i < len(s); i++ {
		fmt.Println(i)
		if s[i] == elem {
			s = append(s[:i], s[i+1:]...)
			i--
		}
	}
	return s
}

func main() {
	a := strings.Split("",",")
	fmt.Println(a,len(a))
	s := []int{1, 2, 3, 4, 5}
	fmt.Println(len(s))
	fmt.Println(cap(s))
	fmt.Println(s[5:])
	fmt.Println(DeleteSlice(s, 5))
}
