package greedy

/**
给定字符串 s 和 t ，判断 s 是否为 t 的子序列。你可以认为 s 和 t 中仅包含英文小写字母。
字符串 t 可能会很长（长度 ~= 500,000），而 s 是个短字符串（长度 <=100）。
字符串的一个子序列是原始字符串删除一些（也可以不删除）字符而不改变剩余字符相对位置形成的新字符串。
（例如，“ace"是"abcde"的一个子序列，而"aec"不是）。
*/

// O(n)
func IsSubSequence(s, t string) bool {
	for len(s) > 0 && len(t) > 0 {
		if s[0] == t[0] {
			s = s[1:]
		}
		t = t[1:]
	}
	return len(s) == 0
}

// O(n^2)
func IsSubSequence2(s, t string) bool {
	j := 0
	for i := 0; i < len(s); i++ {
		flag := false
		for ; j < len(t); j++ {
			if s[i]==t[j]{
				flag = true
				break
			}
		}
		if flag {
			j++
			continue
		}else{
			return false
		}
	}
	return true
}
