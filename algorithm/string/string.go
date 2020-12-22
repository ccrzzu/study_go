package string

import (
	"strconv"
	"strings"
)

//备注：字符串匹配算法
//KMP、SUNDAY、BF等三种

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

//字符串旋转后是否一样
func rotateStringByContain(A string, B string) bool {
	return len(A) == len(B) && strings.Contains(A+A, B)
}

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

//大数相加，两个字符串数字相加，返回string结果
func AddStrings(num1 string, num2 string) string {
	len1, len2 := len(num1), len(num2)
	var result string
	var carry, digit int
	for i, j := len1-1, len2-1; i >= 0 || j >= 0 || carry > 0; {
		digit = carry
		if i >= 0 {
			digit += int(num1[i] - '0')
			i--
		}
		if j >= 0 {
			digit += int(num2[j] - '0')
			j--
		}
		if digit >= 10 {
			carry = 1
		} else {
			carry = 0
		}
		result = strconv.Itoa(digit%10) + result
	}
	return result
}
