package array

import "strings"

/**
给定一个单词列表，只返回可以使用在键盘同一行的字母打印出来的单词。
*/

func findWordsInKeyboardRow(words []string) []string {
	res := []string{}
	keyboradRows := []string{"qwertyuiop", "asdfghjkl", "zxcvbnm"}
	for _, w := range words {
		if len(w) == 0 {
			continue
		}
		lowerW := strings.ToLower(w)
		flag := false
		for _, item := range keyboradRows {
			if strings.ContainsAny(lowerW, item) {
				flag = !flag
				if !flag {
					break
				}
			}
		}
		if flag {
			res = append(res, w)
		}
	}
	return res
}
