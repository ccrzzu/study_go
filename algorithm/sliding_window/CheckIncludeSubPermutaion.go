package sliding_window

/**
567
给你两个字符串 s1 和 s2 ，写一个函数来判断 s2 是否包含 s1 的排列。如果是，返回 true ；否则，返回 false 。
换句话说，s1 的排列之一是 s2 的 子串 。
s1 和 s2 仅包含小写字母.
输入：s1 = "ab" s2 = "eidbaooo"
输出：true
解释：s2 包含 s1 的排列之一 ("ba").
*/
func CheckInclusion(s1 string, s2 string) bool {
	window := map[byte]int{}
	need := map[byte]int{}
	for i := range s1 {
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

func CheckInclusion2(s1 string, s2 string) bool {
	var freq [128]int
	if len(s2) == 0 || len(s2) < len(s1) {
		return false
	}
	for i := 0; i < len(s1); i++ {
		freq[s1[i]-'a']++
	}
	left, right, count := 0, 0, len(s1)

	for right < len(s2) {
		if freq[s2[right]-'a'] >= 1 {
			count--
		}
		freq[s2[right]-'a']--
		right++
		if count == 0 {
			return true
		}
		if right-left == len(s1) {
			if freq[s2[left]-'a'] >= 0 {
				count++
			}
			freq[s2[left]-'a']++
			left++
		}
	}
	return false
}
