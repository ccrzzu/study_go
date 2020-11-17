package algorithm

import "strings"

func removeKdigits(num string, k int) string {
	stackSlice := []byte{}
	for i := range num {
		digit := num[i]
		for k > 0 && len(stackSlice) > 0 && digit < stackSlice[len(stackSlice)-1] {
			stackSlice = stackSlice[:len(stackSlice)-1]
			k--
		}
		stackSlice = append(stackSlice, digit)
	}
	//如果我们删除了 mm 个数字且 m<km<k，这种情况下我们需要从序列尾部删除额外的 k-mk−m 个数字。
	stackSlice = stackSlice[:len(stackSlice)-k]
	//删除所有前导0
	ans := strings.TrimLeft(string(stackSlice), "0")
	if ans == ""{
		ans = "0"
	}
	return ans
}
