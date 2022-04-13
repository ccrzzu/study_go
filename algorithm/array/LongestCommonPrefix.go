package array

import "strings"

//字符串数组的最长公共前缀
//查找字符串数组中的最长公共前缀。如果不存在公共前缀，则返回""
func LongestCommonPrefix(strArr []string) string {
	if len(strArr) == 0 {
		return ""
	}
	prefix := strArr[0]
	for _, item := range strArr {
		for strings.Index(item, prefix) != 0 {
			if prefix == "" {
				return ""
			}
			prefix = prefix[:len(prefix)-1]
		}
	}
	return prefix
}

func LongestCommonPrefix2(strArr []string) string {
	if len(strArr) == 0 {
		return ""
	}
	prefix := strArr[0]
	for i:=1;i<len(strArr);i++{
		prefix = lcp(prefix,strArr[i])
		if prefix == ""{
			break
		}
	}
	return prefix
}

func lcp(str1, str2 string) string {
	length := min(len(str1), len(str2))
	index := 0
	for index < length && str1[index] == str2[index] {
		index++
	}
	return str1[0:index]
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}


func longestCommonPrefix(strArr []string) string {
    if len(strArr) == 0 {
        return ""
    }
    for i := 0; i < len(strArr[0]); i++ {
        for j := 1; j < len(strArr); j++ {
            // if i == len(strArr[j]) || strArr[j][i] != strArr[0][i] {
            //     return strArr[0][:i]
            // }
			if len(strArr[j]) > i && strArr[j][i] == strArr[0][i]{
				continue
			} else {
				return strArr[0][0:i]
			}
        }
    }
    return strArr[0]
}