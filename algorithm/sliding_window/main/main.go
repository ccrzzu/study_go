package main

import (
	"fmt"
	//"study_go/algorithm/sliding_window"
	str "study_go/algorithm/string"
)

func main() {
	// a := sliding_window.LengthOfLongestSubstring4("pwwkew")
	// fmt.Println(a)
	b := str.FindAnagrams2("cbaebabacd", "abc")
	fmt.Println(b)
}
