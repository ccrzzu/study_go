package dynamic_programming

import (
	"fmt"
	"math"
)

/**
121 买卖股票的最佳时机
给定一个数组 prices ，它的第 i 个元素 prices[i] 表示一支给定股票第 i 天的价格。
你只能选择 某一天 买入这只股票，并选择在 未来的某一个不同的日子 卖出该股票。
设计一个算法来计算你所能获取的最大利润。

返回你可以从这笔交易中获取的最大利润。如果你不能获取任何利润，返回 0 。
*/
//dpi0 交易完后手里没有股票的最大利润
//dpi1 交易完后手里持有一支股票的最大利润
func maxProfit(prices []int) int {
	dpi0, dpi1 := 0, math.MinInt32
	for i := 0; i < len(prices); i++ {
		dpi0 = max(dpi0, dpi1+prices[i])
		dpi1 = max(dpi1, -prices[i])
	}
	return dpi0
}

// 解法一 dp
func maxProfit2(prices []int) int {
	if len(prices) < 1 {
		return 0
	}
	min, maxProfit := prices[0], 0
	for i := 1; i < len(prices); i++ {
		if prices[i]-min > maxProfit {
			maxProfit = prices[i] - min
		}
		if prices[i] < min {
			min = prices[i]
		}
	}
	return maxProfit
}

// 解法二 构建单调递增栈，然后由栈顶-栈底即可得出最大利润值
func MaxProfit3(prices []int) int {
	if len(prices) == 0 {
		return 0
	}
	stack, res := []int{prices[0]}, 0
	for i := 1; i < len(prices); i++ {
		if prices[i] > stack[len(stack)-1] {
			stack = append(stack, prices[i])
		} else {
			index := len(stack) - 1
			for ; index >= 0; index-- {
				if stack[index] < prices[i] {
					break
				}
			}
			fmt.Println(stack, index)
			stack = stack[:index+1]
			fmt.Println(stack)
			stack = append(stack, prices[i])
		}
		fmt.Println(stack)
		res = max(res, stack[len(stack)-1]-stack[0])
	}
	return res
}

//122
//与121不同的是，可以多次交易
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

func maxProfit_k_inf2(prices []int) (ans int) {
    for i := 1; i < len(prices); i++ {
        ans += max(0, prices[i]-prices[i-1])
    }
    return
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
