package dynamic_programming

import "math"

//给定一个数组，它的第 i 个元素是一支给定股票第 i 天的价格。
//如果你最多只允许完成一笔交易（即买入和卖出一支股票一次），设计一个算法来计算你所能获取的最大利润。
func maxProfit(prices []int) int {
	dpi0, dpi1 := 0, math.MinInt32
	for i := 0; i < len(prices); i++ {
		dpi0 = int(math.Max(float64(dpi0), float64(dpi1+prices[i])))
		dpi1 = int(math.Max(float64(dpi1), float64(-prices[i])))
	}
	return dpi0
}

//给定一个整数数组prices ，它的第 i 个元素prices[i] 是一支给定的股票在第 i 天的价格。
//设计一个算法来计算你所能获取的最大利润。你可以尽可能多的完成交易。
func maxProfit_k_inf(prices []int) int {
	dpi0, dpi1 := 0, math.MinInt32
	for i := 0; i < len(prices); i++ {
		//记录上一天的持有
		temp := dpi0
		dpi0 = int(math.Max(float64(dpi0), float64(dpi1+prices[i])))
		dpi1 = int(math.Max(float64(dpi1), float64(temp-prices[i])))
	}
	return dpi0
}

func MaxProfit_k_2(prices []int) int {
	max_k := 2
	dp := make([][][]int, len(prices))
	for i := 0; i < len(prices); i++ {
		dp[i] = make([][]int, max_k+1)
		for j := 0; j < max_k+1; j++ {
			dp[i][j] = make([]int, 2)
		}
	}
	for i := 0; i < len(prices); i++ {
		for k := max_k; k >= 1; k-- {
			if i-1 == -1 {
				dp[i][k][0] = 0
				dp[i][k][1] = -prices[i]
				continue
			}
			dp[i][k][0] = int(math.Max(float64(dp[i-1][k][0]), float64(dp[i-1][k][1]+prices[i])))
			dp[i][k][1] = int(math.Max(float64(dp[i-1][k][1]), float64(dp[i-1][k-1][0]-prices[i])))
		}
	}
	return dp[len(prices)-1][max_k][0]
}

func MaxProfit_k_s(maxK int, prices []int) int {
	if maxK > len(prices)/2 {
		return maxProfit_k_inf(prices)
	}
	dp := make([][][]int, len(prices))
	for i := 0; i < len(prices); i++ {
		dp[i] = make([][]int, maxK+1)
		for j := 0; j < maxK+1; j++ {
			dp[i][j] = make([]int, 2)
		}
	}
	for i := 0; i < len(prices); i++ {
		for k := maxK; k >= 1; k-- {
			if i-1 == -1 {
				dp[i][k][0] = 0
				dp[i][k][1] = -prices[i]
				continue
			}
			dp[i][k][0] = int(math.Max(float64(dp[i-1][k][0]), float64(dp[i-1][k][1]+prices[i])))
			dp[i][k][1] = int(math.Max(float64(dp[i-1][k][1]), float64(dp[i-1][k-1][0]-prices[i])))
		}
	}
	return dp[len(prices)-1][maxK][0]
}

func maxProfit_with_fee(prices []int, fee int) int {
	dpi0, dpi1 := 0, math.MinInt32
	for i := 0; i < len(prices); i++ {
		//记录上一天的持有
		temp := dpi0
		dpi0 = int(math.Max(float64(dpi0), float64(dpi1+prices[i])))
		dpi1 = int(math.Max(float64(dpi1), float64(temp-prices[i]-fee)))
	}
	return dpi0
}

//t+2买入
func maxProfit_with_cooldown(prices []int) int {
	dpi0, dpi1 := 0, math.MinInt32
	preTemp := 0
	for i := 0; i < len(prices); i++ {
		//记录上一天的持有
		temp := dpi0
		dpi0 = int(math.Max(float64(dpi0), float64(dpi1+prices[i])))
		dpi1 = int(math.Max(float64(dpi1), float64(preTemp-prices[i])))
		preTemp = temp
	}
	return dpi0
}
