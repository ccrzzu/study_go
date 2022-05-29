package dynamic_programming

import "math"

//给定两个字符串 text1 和 text2，返回这两个字符串的最长公共子序列的长度。
func LongestCommonSubsequence(text1 string, text2 string) int {
	//return dgForLCS(text1, text2, len(text1)-1, len(text2)-1)
	t1l := len(text1)
	t2l := len(text2)
	dp := make([][]int, t1l+1)
	//构建dp table和base case
	for i := 0; i <= t1l; i++ {
		dp[i] = make([]int, t2l+1)
	}
	for i := 1; i <= t1l; i++ {
		for j := 1; j <= t2l; j++ {
			//找到一个都在lcs中的字符
			if text1[i-1] == text2[j-1] {
				dp[i][j] = dp[i-1][j-1] + 1
			} else {
				dp[i][j] = int(math.Max(float64(dp[i][j-1]), float64(dp[i-1][j])))
			}
		}
	}
	return dp[t1l][t2l]
}

//最长公共子序列的递归解法
func dgForLCS(text1, text2 string, i, j int) int {
	//当都是空串的时候，返回-1
	if i == -1 || j == -1 {
		return 0
	}
	if text1[i] == text2[j] {
		//找到一个lcs的元素，go on search
		return dgForLCS(text1, text2, i-1, j-1) + 1
	} else {
		//谁让lcs最长，就选谁
		return int(math.Max(float64(dgForLCS(text1, text2, i-1, j)), float64(dgForLCS(text1, text2, i, j-1))))
	}
}
