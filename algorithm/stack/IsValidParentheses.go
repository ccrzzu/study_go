package stack

/**
给定一个仅包含字符 '(', ')', '{', '}', '[' 和 ']' 的字符串，确定输入字符串是否有效。

输入字符串在以下情况下有效：

开括号必须用相同类型的括号闭合。开括号必须以正确的顺序闭合。请注意，空字符串也被认为是有效的。
*/

func IsValidParentheses(s string) bool {
	if len(s) == 0 {
		return true
	}
	stack := make([]rune, 0)
	for _, item := range s {
		if item == '(' || item == '{' || item == '[' {
			stack = append(stack, item)
		} else if (item == ')' && len(stack) > 0 && stack[len(stack)-1] == '(') ||
			(item == '}' && len(stack) > 0 && stack[len(stack)-1] == '{') ||
			(item == ']' && len(stack) > 0 && stack[len(stack)-1] == '[') {
			stack = stack[:len(stack)-1]
		} else {
			return false
		}
	}
	return len(stack)==0
}
