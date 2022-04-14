package string

//字符串中的第一个唯一字符
func firstUniqChar(s string) int {
	m := [26]int{}
	for i := 0; i < len(s); i++ {
		m[s[i]-'a'] = i
	}

	for i, item := range s {
		if m[item-'a'] == i {
			return i
		} else {
			m[item-'a'] = -1
		}
	}
	return -1
}