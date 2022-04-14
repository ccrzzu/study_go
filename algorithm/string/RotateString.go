package string

import "strings"

//字符串旋转后是否一样
func rotateString(A string, B string) bool {
	if A == "" && B == "" {
		return true
	}
	for i := 0; i < len(A); i++ {
		temp := string(A[i+1:]) + string(A[:i+1])
		if temp == B {
			return true
		}
	}
	return false
}

//字符串旋转后是否一样 面试不能这样写 
func rotateStringByContain(A string, B string) bool {
	return len(A) == len(B) && strings.Contains(A+A, B)
}