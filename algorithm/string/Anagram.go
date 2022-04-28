package string

/**
242 判断两个字符串是否是彼此的相同字母异序词
给定两个字符串 s 和 t ，编写一个函数来判断 t 是否是 s 的字母异位词。
注意：若 s 和 t 中每个字符出现的次数都相同，则称 s 和 t 互为字母异位词。
注意字母大小写，是否题意被忽略。
*/
func ValidAnagram1(s, t string) bool {
	res := make([]byte, 26)
	for i := 0; i < len(s); i++ {
		res[s[i]-'a']++
	}
	for i := 0; i < len(t); i++ {
		res[t[i]-'a']--
	}
	for _, item := range res {
		if item != 0 {
			return false
		}
	}
	return true
}

func ValidAnagram2(s, t string) bool {
	res := make(map[rune]int)
	for _, sItem := range s {
		res[sItem]++
	}
	for _, tItem := range t {
		res[tItem]--
	}
	for _, item := range res {
		if item != 0 {
			return false
		}
	}
	return true
}
