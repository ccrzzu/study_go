package algorithm

import "strings"

//总述：
//贪心算法，又名贪婪法，是寻找最优解问题的常用方法，
//这种方法模式一般将求解过程分成若干个步骤，但每个步骤都应用贪心原则，
//选取当前状态下最好/最优的选择（局部最有利的选择），
//并以此希望最后堆叠出的结果也是最好/最优的解。

//给定一个以字符串表示的非负整数 num，移除这个数中的 k 位数字，使得剩下的数字最小。
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
	if ans == "" {
		ans = "0"
	}
	return ans
}

//在一条环路上有 N 个加油站，其中第 i 个加油站有汽油 gas[i] 升。
//你有一辆油箱容量无限的的汽车，从第 i 个加油站开往第 i+1 个加油站需要消耗汽油 cost[i] 升。你从其中的一个加油站出发，开始时油箱为空。
//如果你可以绕环路行驶一周，则返回出发时加油站的编号，否则返回 -1。
func canCompleteCircuit(gas []int, cost []int) int {
	left, start := 0, 0
	totalCost, totalGas := 0, 0
	for i := 0; i < len(gas); i++ {
		left += gas[i] - cost[i]
		totalCost += cost[i]
		totalGas += gas[i]
		if left < 0 {
			start = i + 1
			left = 0
		}
	}
	if totalGas < totalCost {
		return -1
	}
	return start
}
