package dynamic_programming

/**
 * 本节是爬楼梯或者叫跳台阶等问题的集合
 * 经典的斐波那契数列问题
 */

/*
 * 暴力递归的升级版：递归的优化解法，增加了记忆，将时间复杂度降低至O(n)
 */
func ClimbStairsByRecur(n int) int {
	memo := map[int]int{0: 1, 1: 1, 2: 2}
	var climbStairsByRecur func(n int) int
	climbStairsByRecur = func(n int) int {
		if res, ok := memo[n]; ok {
			//fmt.Println(n, memo[n])
			return res
		}
		memo[n] = climbStairsByRecur(n-1) + climbStairsByRecur(n-2)
		return memo[n]
	}
	return climbStairsByRecur(n)
}

/*
 * 动态规划解法，找到最优子结构，列出动态转移方程dp[i] = dp[i - 1] + dp[ i - 2];
 */
func ClimbStairsByDynamicProgram(n int) int {
	/*
		dp := map[int]int{0: 1, 1: 1, 2: 2}
		for i := 3; i <= n; i++ {
			dp[i] = dp[i-1] + dp[i-2]
		}
		return dp[n]
	*/
	//dp := []int{0, 1, 2}
	//for i := 3; i <= n; i++ {
	//	dp = append(dp, 0)
	//}
	dp := make([]int, n+1)
	dp[0] = 1
	if n >= 1 {
		dp[1] = 1
	}
	if n >= 2 {
		dp[2] = 2
	}
	for i := 3; i <= n; i++ {
		dp[i] = dp[i-1] + dp[i-2]
	}
	return dp[n]
}

/**
 * 由于机器总线原因，整型值有最大值一说
 * 此为需要除余的，答案需要取模 1e9+7（1000000007）
 */
func ClimbStairsForDivideAndRemainder(n int) int {
	memo := map[int]int{0: 1, 1: 1, 2: 2}
	var climbStairsByRecur func(n int) int
	climbStairsByRecur = func(n int) int {
		if res, ok := memo[n]; ok {
			//fmt.Println(n, memo[n])
			return res
		}
		memo[n] = (climbStairsByRecur(n-1) + climbStairsByRecur(n-2)) % (1e9 + 7)
		return memo[n]
	}
	return climbStairsByRecur(n)
}

/**
 * 除余版 动态规划解法
 */
func ClimbStairsForDivideAndRemainder2(n int) int {
	dp := map[int]int{0: 1, 1: 1, 2: 2}
	for i := 3; i <= n; i++ {
		dp[i] = (dp[i-1] + dp[i-2]) % (1E9 + 7)
	}
	return dp[n]
}

/**
 * 三步问题 就是上面两步问题的延伸 思维是一样的 就很简单了
 */
func WaysToStepByArray(n int) int {
	//动态规划版解法 by array
	dp := []int{0: 1, 1: 1, 2: 2, 3: 4}
	for i := 4; i <= n; i++ {
		dp = append(dp, 0)
	}
	for i := 3; i <= n; i++ {
		dp[i] = (dp[i-1] + dp[i-2] + dp[i-3]) % (1E9 + 7)
	}
	return dp[n]
}

/**
 * 动态规划解法 by map
 */
func WaysToStepByMap(n int) int {
	//动态规划版解法 by map
	dp := map[int]int{0: 1, 1: 1, 2: 2, 3: 4}
	for i := 3; i <= n; i++ {
		dp[i] = (dp[i-1] + dp[i-2] + dp[i-3]) % (1E9 + 7)
	}
	return dp[n]
}
