package sliding_window

import "fmt"

/**
 438
 给定两个字符串 s 和 p，找到 s 中所有 p 的 异位词 的子串，返回这些子串的起始索引。不考虑答案输出的顺序。
异位词 指由相同字母重排列形成的字符串（包括相同的字符串）。
输入: s = "abab", p = "ab"
输出: [0,1,2]
解释:
起始索引等于 0 的子串是 "ab", 它是 "ab" 的异位词。
起始索引等于 1 的子串是 "ba", 它是 "ab" 的异位词。
起始索引等于 2 的子串是 "ab", 它是 "ab" 的异位词。
*/
func FindAnagrams1(s string, p string) []int {
	window := map[byte]int{}
	need := map[byte]int{}
	for i := range p {
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

func FindAnagrams2(s string, p string) []int {
	var freq [128]int
	result := []int{}
	if len(s) == 0 || len(s) < len(p) {
		return result
	}
	for i := 0; i < len(p); i++ {
		freq[p[i]-'a']++
	}
	left, right, count := 0, 0, len(p)
	
	for right < len(s) {
		fmt.Println(freq, count)
		if freq[s[right]-'a'] >= 1 {
			count--
		}
		freq[s[right]-'a']--
		right++
		if count == 0 {
			result = append(result, left)
		}
		if right-left == len(p) {
			if freq[s[left]-'a'] >= 0 {
				count++
			}
			freq[s[left]-'a']++
			left++
		}
		
	}
	return result
}
