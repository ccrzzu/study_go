package string

import "strings"

//最后一个单词的长度 偷懒
func lengthOfLastWordCrazy(s string) int {
	s = strings.TrimSpace(s)
	index := strings.LastIndex(s, " ")
	return len(s[index+1:])
}

func lengthOfLastWordCrazy2(s string) int {
	s = strings.TrimSpace(s)
	split := strings.Split(s, " ")
	if len(split) < 1 {
		return 0
	}
	return len(split[len(split)-1])
}

//最后一个单词的长度
func lengthOfLastWordByCount(s string) int {
	if len(s) == 0 {
		return 0
	}
	count := 0
	for i := len(s) - 1; i >= 0; i-- {
		if string(s[i]) == " " {
			if count == 0 {
				continue
			}
			break
		}
		count++
	}
	return count
}