package main

import (
	"fmt"
	"study_go/algorithm/sliding_window"
)

func main() {
	// a := sliding_window.LengthOfLongestSubstring4("pwwkew")
	// fmt.Println(a)
	b := sliding_window.FindAnagrams2("cbaebabacd", "abc")
	fmt.Println(b)
}
