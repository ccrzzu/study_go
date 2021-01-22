package stack

import (
	"math"
	"strings"
)

//输入两个整数序列，第一个序列表示栈的压入顺序，请判断第二个序列是否可能为该栈的弹出顺序。
//假设压入栈的所有数字均不相等。例如序列1,2,3,4,5是某栈的压入顺序，
//序列4,5,3,2,1是该压栈序列对应的一个弹出序列，但4,3,5,1,2就不可能是该压栈序列的弹出序列。
//（注意：这两个序列的长度是相等的）
func IsPopOrder(pushed []int, popped []int) bool {
	stack := []int{}
	i, j := 0, 0
	for i < len(pushed) && j < len(popped) {
		stack = append(stack, pushed[i])
		for len(stack) > 0 && stack[len(stack)-1] == popped[j] {
			j++
			stack = stack[:len(stack)-1]
		}
		i++
	}
	return len(stack) == 0
}

func IsPopOrder2(pushed []int, popped []int) bool {
	stack := []int{}
	for i := 0; i < len(pushed); i++ {
		stack = append(stack, pushed[i])
		for len(stack) > 0 && popped[0] == stack[len(stack)-1] {
			popped = popped[1:]
			stack = stack[:len(stack)-1]
		}
	}
	return len(stack) == 0
	//return len(popped) == 0
}

//删除一个字符串的对称的最外层括号
func removeOuterParentheses(S string) string {
	flag := 0
	b := 1 //从第几个位开始作为内层算起
	var res string
	for i := 0; i < len(S); i++ {
		if S[i] == '(' {
			flag++
		} else {
			flag--
		}
		if flag == 0 {
			res += S[b:i]
			b = i + 2
		}
	}
	return res
}

func removeOuterParentheses2(S string) string {
	if len(S) == 0 {
		return ""
	}
	stack := []byte{}
	start := 1
	var res strings.Builder
	for i := 0; i < len(S); i++ {
		if len(stack) > 0 && S[i] == ')' && stack[len(stack)-1] == '(' {
			stack = stack[:len(stack)-1]
			if len(stack) == 0 {
				res.WriteString(S[start:i])
				start = i + 2
			}
		} else {
			stack = append(stack, S[i])
		}
	}
	return res.String()
}

//移除字符串中所有相邻的
func removeDuplicates(S string) string {
	if len(S) == 0 {
		return ""
	}
	stack := []byte{}
	for i := 0; i < len(S); i++ {
		stack = append(stack, S[i])
		for len(stack) >= 2 && stack[len(stack)-1] == stack[len(stack)-2] {
			stack = stack[:len(stack)-2]
		}
	}
	return string(stack)
}

func removeDuplicates2(S string) string {
	stack := []byte{}
	for i := 0; i < len(S); i++ {
		if len(stack) > 0 && S[i] == stack[len(stack)-1] {
			stack = stack[:len(stack)-1]
		} else {
			stack = append(stack, S[i])
		}
	}
	return string(stack)
}

/*
[3 1 2 4]
num stack   stacksum  sum     subarray
3   3,1            3   3      [3]
1   1,2            2   5      [3 1]     [1]
2   1,2 2,1        4   9      [3 1 2]   [1 2]   [2]
4   1,2 2,1 4,1    8   17     [3 1 2 4] [1 2 4] [2 4] [4]
栈[num,count]代表子数组最小值，以及数组个数
*/
func SumSubarrayMins(A []int) int{
	type pair struct {
		value int
		count int
	}
	stack := make([]pair, 0)
	var sum,stackSum int
	div := int(math.Pow(10, 9) + 7)
	for _, num := range A {
		count := 1
		for len(stack) != 0 && num < stack[len(stack)-1].value {
			pop := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			stackSum -= pop.value*pop.count
			count += pop.count
		}
		node := pair{num, count}
		stack = append(stack, node)
		stackSum += node.value*node.count
		sum += stackSum
		//sum %= div
	}
	return sum % div
}
