package backtrack

import "fmt"

/**
22
数字 n 代表生成括号的对数，请你设计一个函数，用于能够生成所有可能的并且 有效的 括号组合。
*/

func GenerateParenthesis(n int) []string {
	res := []string{}

	var dfs func(lRemain int, rRemain int, path string)
	dfs = func(lRemain int, rRemain int, path string) {
		if 2*n == len(path) {
			res = append(res, path)
			fmt.Println("result:", res)
			return
		}
		if lRemain > 0 {
			fmt.Printf("left lRemain:%d, rRemain:%d ,res:%s\n", lRemain, rRemain, path)
			dfs(lRemain-1, rRemain, path+"(")
		}
		if lRemain < rRemain {
			fmt.Printf("right lRemain:%d, rRemain:%d ,res:%s\n", lRemain, rRemain, path)
			dfs(lRemain, rRemain-1, path+")")
		}
	}

	dfs(n, n, "")
	return res
}
