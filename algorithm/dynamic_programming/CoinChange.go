package dynamic_programming

import (
	"fmt"
	"math"
)

/**
 * @param coins: a list of integer
 * @param amount: a total amount of money amount
 * @return: the fewest number of coins that you need to make up
 */
//凑零钱 递归
func CoinChange(coins []int, amount int) int {
	// write your code here
	memo := make(map[int]int)
	//fmt.Println(math.MaxInt64)
	coinsMap := make(map[int]bool)
	for _, item := range coins {
		if item == 0 {
			continue
		}
		coinsMap[item] = true
	}
	newCoins := make([]int, 0)
	for k := range coinsMap {
		newCoins = append(newCoins, k)
	}
	return dg(memo, newCoins, amount)
}

func dg(memo map[int]int, coins []int, amount int) int {
	// 增加备忘录，避免重复计算
	if i, ok := memo[amount]; ok {
		return i
	}
	// base case
	if amount == 0 {
		return 0
	}
	if amount < 0 {
		return -1
	}
	result := math.MaxInt64
	for _, coin := range coins {
		subProblem := dg(memo, coins, amount-coin)
		if subProblem == -1 {
			continue
		}
		result = int(math.Min(float64(result), float64(1+subProblem)))
	}

	// 将结果写入备忘录
	if result != math.MaxInt64 {
		memo[amount] = result
	} else {
		memo[amount] = -1
	}
	fmt.Printf("memo[%v],%v\n", amount, result)
	return memo[amount]
}

//凑零钱 迭代
func CoinChangeByIter(coins []int, amount int) int {
	// 设定数组大小为amount+1，同时初始值也都设为amount+1
	amountInit := make([]int, amount+1)
	for i := 0; i < len(amountInit); i++ {
		amountInit[i] = i + 1
	}
	// base case
	amountInit[0] = 0
	// 外层循环遍历所有状态的所有取值
	for i := 0; i < len(amountInit); i++ {
		// 内层循环求所有选择的最小值
		for _, coin := range coins {
			// 小于0的肯定是无解的，跳过
			if i-coin < 0 {
				continue
			}
			if amountInit[i-coin] == -1 {
				continue
			}
			amountInit[i] = int(math.Min(float64(amountInit[i]), float64(1+amountInit[i-coin])))
		}
		if amountInit[i] == i+1 {
			amountInit[i] = -1
		}
	}
	return amountInit[amount]
}
