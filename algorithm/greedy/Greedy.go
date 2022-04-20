package greedy

//总述：
//贪心算法，又名贪婪法，是寻找最优解问题的常用方法，
//这种方法模式一般将求解过程分成若干个步骤，但每个步骤都应用贪心原则，
//选取当前状态下最好/最优的选择（局部最有利的选择），
//并以此希望最后堆叠出的结果也是最好/最优的解。

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
