package stack

/**
 * 删除字符串中所有相邻且相同的字母
 *给出由小写字母组成的字符串 S，重复项删除操作会选择两个相邻且相同的字母，并删除它们。
 *在 S 上反复执行重复项删除操作，直到无法继续删除。
 *在完成所有重复项删除操作后返回最终的字符串。答案保证唯一。
 */
func RemoveDuplicates(S string) string {
	if len(S) == 0 {
		return ""
	}
	stack := []byte{}
	for i := 0; i < len(S); i++ {
		stack = append(stack, S[i])
		for len(stack) >= 2 && stack[len(stack)-1] == stack[len(stack)-2] {
			stack = stack[:len(stack)-2]
		}
	}
	return string(stack)
}

func RemoveDuplicates2(S string) string {
	stack := []byte{}
	for i := 0; i < len(S); i++ {
		if len(stack) > 0 && S[i] == stack[len(stack)-1] {
			stack = stack[:len(stack)-1]
		} else {
			stack = append(stack, S[i])
		}
	}
	return string(stack)
}