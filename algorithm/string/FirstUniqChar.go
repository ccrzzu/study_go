package string

//字符串中的第一个唯一字符
//给定一个字符串，找到它的第一个不重复的字符，并返回它的索引。如果不存在，则返回 -1。
func firstUniqChar(s string) int {
	m := [26]int{}
	for _, item := range s {
		m[item-'a']++
	}
	for i, item := range s {
		if m[item-'a'] == 1 {
			return i
		}
	}
	return -1
}

func firstUniqChar2(s string) int {
	charMap := make([][2]int, 26)
	for i := 0; i < 26; i++ {
		charMap[i][0] = -1
		charMap[i][1] = -1
	}
	for i := 0; i < len(s); i++ {
		if charMap[s[i]-'a'][0] == -1 {
			charMap[s[i]-'a'][0] = i
		} else {
			charMap[s[i]-'a'][1] = i
		}
	}
	res := len(s)
	for i := 0; i < 26; i++ {
		if charMap[i][0] >= 0 && charMap[i][1] == -1 {
			if charMap[i][0] < res{
				res = charMap[i][0]
			}
		}
	}
	if res == len(s) {
		return -1
	}
	return res
}

func firstUniqChar3(s string) int {
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
