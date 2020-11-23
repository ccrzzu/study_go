package double_pointer

import "math"

func minWindow(s string, t string) string {
	window := map[byte]int{}
	need := map[byte]int{}
	for i, _ := range t {
		need[t[i]]++
	}
	left, right, valid := 0, 0, 0
	start, length := 0, math.MaxInt32
	for right < len(s) {
		rightChar := s[right]
		right++
		if _, ok := need[rightChar]; ok {
			window[rightChar]++
			if need[rightChar] == window[rightChar] {
				valid++
			}
		}

		for valid == len(need) {
			if right-left < length {
				start = left
				length = right - left
			}
			leftChar := s[left]
			left++
			if _, ok := need[leftChar]; ok {
				if need[leftChar] == window[leftChar] {
					valid--
				}
				window[leftChar]--
			}
		}
	}
	if length != math.MaxInt32 {
		return s[start : start+length]
	}
	return ""
}

func checkInclusion(s1 string, s2 string) bool {
	window := map[byte]int{}
	need := map[byte]int{}
	for i, _ := range s1 {
		need[s1[i]]++
	}
	left, right, valid := 0, 0, 0
	for right < len(s2) {
		rightChar := s2[right]
		right++
		if _, ok := need[rightChar]; ok {
			window[rightChar]++
			if need[rightChar] == window[rightChar] {
				valid++
			}
		}

		for right-left == len(s1) {
			if valid == len(need) {
				return true
			}
			leftChar := s2[left]
			left++
			if _, ok := need[leftChar]; ok {
				if need[leftChar] == window[leftChar] {
					valid--
				}
				window[leftChar]--
			}
		}
	}
	return false
}

func findAnagrams(s string, p string) []int {
	window := map[byte]int{}
	need := map[byte]int{}
	for i, _ := range p {
		need[p[i]]++
	}
	left, right, valid := 0, 0, 0
	res := []int{}
	for right < len(s) {
		rightChar := s[right]
		right++
		if _, ok := need[rightChar]; ok {
			window[rightChar]++
			if need[rightChar] == window[rightChar] {
				valid++
			}
		}

		for right-left == len(p) {
			if valid == len(need) {
				res = append(res, left)
			}
			leftChar := s[left]
			left++
			if _, ok := need[leftChar]; ok {
				if need[leftChar] == window[leftChar] {
					valid--
				}
				window[leftChar]--
			}
		}
	}
	return res
}

func lengthOfLongestSubstring(s string) int {
	window := map[byte]int{}
	left,right,res := 0,0,0
	for right < len(s){
		rightChar := s[right]
		right++
		window[rightChar]++
		for window[rightChar] > 1{
			leftChar := s[left]
			left++
			window[leftChar]--
		}
		res = int(math.Max(float64(res),float64(right -left)))
	}
	return res
}
