package string

/**
给定两个字符串 s 和 t，它们只包含小写字母。
字符串 t 由字符串 s 随机重排，然后在随机位置添加一个字母。请找出在 t 中被添加的字母。
思路：所有字母参与异或^操作，一样的归为0，不一样的为1，
举例： "abc" "caeb"  b^a^c^b^a^c^e 相同的消掉后剩的就是多出来的那个e
*/
func findTheDifference(s string, t string) byte {
	res := t[len(t)-1]
	for i := 0; i < len(s); i++ {
		res ^= s[i]
		res ^= t[i]
	}
	return res
}
