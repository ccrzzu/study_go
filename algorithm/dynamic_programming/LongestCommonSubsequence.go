package dynamic_programming

import "math"

/**
1143
给定两个字符串 text1 和 text2，返回这两个字符串的最长 公共子序列 的长度。
如果不存在 公共子序列 ，返回 0 。

一个字符串的 子序列 是指这样一个新的字符串：
它是由原字符串在不改变字符的相对顺序的情况下删除某些字符（也可以不删除任何字符）后组成的新字符串。

例如，"ace" 是 "abcde" 的子序列，但 "aec" 不是 "abcde" 的子序列。
两个字符串的 公共子序列 是这两个字符串所共同拥有的子序列。
*/

//解法一 动态规划解法
func LongestCommonSubsequence(text1 string, text2 string) int {
	//递归解法写法：
	//return LongestCommonSubsequenceByDG(text1, text2, len(text1)-1, len(text2)-1)
	
	t1len := len(text1)
	t2len := len(text2)
	dp := make([][]int, t1len+1)
	//构建dp table和base case
	for i := 0; i <= t1len; i++ {
		dp[i] = make([]int, t2len+1)
	}
	for i := 1; i <= t1len; i++ {
		for j := 1; j <= t2len; j++ {
			//找到一个都在lcs中的字符
			if text1[i-1] == text2[j-1] {
				dp[i][j] = dp[i-1][j-1] + 1
			} else {
				dp[i][j] = int(math.Max(float64(dp[i][j-1]), float64(dp[i-1][j])))
			}
		}
	}
	return dp[t1len][t2len]
}

// 解法二：递归解法
func LongestCommonSubsequenceByDG(text1, text2 string, i, j int) int {
	//当都是空串的时候，返回-1
	if i == -1 || j == -1 {
		return 0
	}
	if text1[i] == text2[j] {
		//找到一个lcs的元素，go on search
		return LongestCommonSubsequenceByDG(text1, text2, i-1, j-1) + 1
	} else {
		//谁让lcs最长，就选谁
		return int(math.Max(float64(LongestCommonSubsequenceByDG(text1, text2, i-1, j)),
			float64(LongestCommonSubsequenceByDG(text1, text2, i, j-1))))
	}
}
