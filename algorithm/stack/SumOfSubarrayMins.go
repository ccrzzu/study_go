package stack

import (
	"math"
)

/*
*给定一个整数数组 arr，找到 min(b) 的总和，其中 b 的范围为 arr 的每个（连续）子数组。
由于答案可能很大，因此 返回答案模 10^9 + 7 。

[3 1 2 4]
num stack   stacksum  sum     subarray
3   3,1            3   3      [3]
1   1,2            2   5      [3 1]     [1]
2   1,2 2,1        4   9      [3 1 2]   [1 2]   [2]
4   1,2 2,1 4,1    8   17     [3 1 2 4] [1 2 4] [2 4] [4]
栈[num,count]代表子数组最小值，以及数组个数
*/
func SumOfSubarrayMins(A []int) int {
	type pair struct {
		value int
		count int
	}
	stack := make([]pair, 0)
	var sum, stackSum int
	div := int(math.Pow(10, 9) + 7)
	for _, num := range A {
		count := 1
		for len(stack) != 0 && num < stack[len(stack)-1].value {
			pop := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			stackSum -= pop.value * pop.count
			count += pop.count
		}
		node := pair{num, count}
		stack = append(stack, node)
		stackSum += node.value * node.count
		sum += stackSum
		//sum %= div
	}
	return sum % div
}
