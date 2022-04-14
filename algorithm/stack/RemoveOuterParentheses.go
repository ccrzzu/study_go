package stack

import "strings"

//删除一个字符串的对称的最外层括号
func RemoveOuterParentheses(S string) string {
	flag := 0
	b := 1 //从第几个位开始作为内层算起
	var res string
	for i := 0; i < len(S); i++ {
		if S[i] == '(' {
			flag++
		} else {
			flag--
		}
		if flag == 0 {
			res += S[b:i]
			b = i + 2
		}
	}
	return res
}

func RemoveOuterParentheses2(S string) string {
	if len(S) == 0 {
		return ""
	}
	stack := []byte{}
	start := 1
	var res strings.Builder
	for i := 0; i < len(S); i++ {
		if len(stack) > 0 && S[i] == ')' && stack[len(stack)-1] == '(' {
			stack = stack[:len(stack)-1]
			if len(stack) == 0 {
				res.WriteString(S[start:i])
				start = i + 2
			}
		} else {
			stack = append(stack, S[i])
		}
	}
	return res.String()
}