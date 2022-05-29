package string

import (
	"unicode"
)

/**
1763
当一个字符串 s 包含的每一种字母的大写和小写形式 同时 出现在 s 中，就称这个字符串 s 是 美好 字符串。
比方说，"abABB" 是美好字符串，因为 'A' 和 'a' 同时出现了，且 'B' 和 'b' 也同时出现了。
然而，"abA" 不是美好字符串因为 'b' 出现了，而 'B' 没有出现。

给你一个字符串 s ，请你返回 s 最长的 美好子字符串 。
如果有多个答案，请你返回 最早 出现的一个。如果不存在美好子字符串，请你返回一个空字符串。
*/

//解法一，暴力解法。枚举每一段字符串，判断这个子字符串内是否满足美好字符串的定义，即字母的大小写是否同时出现。
func LongestNiceSubstring(s string) string {
	var res string
	for i := 0; i < len(s); i++ {
		m := map[byte]bool{}
		m[s[i]] = true
		for j := i + 1; j < len(s); j++ {
			m[s[j]] = true
			if checkNiceString(m) && j-i+1 > len(res) {
				res = s[i : j+1]
			}
		}
	}
	return res
}

//检查map的key的byte是否是同时是仅仅存在对应大小写的美好子字符串
func checkNiceString(m map[byte]bool) bool {
	for k := range m {
		//如果该byte是大写26字母之一，则看看有么有不存在的对应的小写字母
		if k >= 65 && k <= 90 {
			if _, ok := m[k+32]; !ok {
				return false
			}
		}
		//如果该byte是小写26字母之一，则看看有么有不存在的对应的大写字母
		if k >= 97 && k <= 122 {
			if _, ok := m[k-32]; !ok {
				return false
			}
		}
	}
	return true
}

// 解法二 用二进制表示状态
func LongestNiceSubstring2(s string) string {
	var res string
	for i := 0; i < len(s); i++ {
		lower, upper := 0, 0
		for j := i; j < len(s); j++ {
			if unicode.IsLower(rune(s[j])) {
				lower |= 1 << (s[j] - 97)
				//lower |= 1 << (s[j] - 'a')
			} else {
				upper |= 1 << (s[j] - 65)
				//upper |= 1 << (s[j] - 'A')
			}
			if lower == upper && j-i+1 > len(res) {
				res = s[i : j+1]
			}
		}
	}
	return res
}

//解法三，分治。以 i 为分割点依次切开字符串。左右两个字符串分别判断是否满足美好字符串的定义。
//左右分开的字符串还可以继续划分。直至分到一个字母为止。在这个过程中记录最早出现的字符串。
func LongestNiceSubstring3(s string) string {
	if len(s) < 2 {
		return ""
	}
	chars := map[rune]int{}
	for _, item := range s {
		chars[item]++
	}
	for i := 0; i < len(s); i++ {
		_, uOk := chars[unicode.ToUpper(rune(s[i]))]
		_, lOk := chars[unicode.ToLower(rune(s[i]))]
		if uOk && lOk {
			continue
		}
		//如果上述逻辑走不通，说明i这个位置的字符肯定不能凑成完美子字符串了，
		//所以以i为分界分治剩下的两个子字符串，分别寻找其完美子字符串
		left := LongestNiceSubstring3(s[:i])
		right := LongestNiceSubstring3(s[i+1:])
		if len(left) >= len(right) {
			return left
		} else {
			return right
		}
	}
	return s
}
