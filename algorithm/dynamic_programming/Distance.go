package dynamic_programming

import (
	"fmt"
	"math"
)

func MinDistance(word1, word2 string) int {
	lw1 := len(word1)
	lw2 := len(word2)
	dp := make([][]int, lw1+1)
	for i := 0; i < lw1+1; i++ {
		dp[i] = make([]int, lw2+1)
	}
	for i := 1; i <= lw1; i++ {
		dp[i][0] = i
	}
	for j := 1; j <= lw2; j++ {
		dp[0][j] = j
	}
	fmt.Println(dp)
	for i := 1; i <= lw1; i++ {
		for j := 1; j <= lw2; j++ {
			if word1[i-1] == word2[j-1] {
				dp[i][j] = dp[i-1][j-1]
			} else {
				dp[i][j] = int(math.Min(math.Min(float64(dp[i-1][j]+1), float64(dp[i][j-1]+1)), float64(dp[i-1][j-1]+1)))
			}
		}
	}
	return dp[lw1][lw2]
}
