package stack

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

func validateStackSequences(pushed []int, popped []int) bool {
	stack := make([]int, 0)
	i, j := 0, 0
	for i < len(pushed) && j < len(popped) {
		stack = append(stack, pushed[i])
		for len(stack) > 0 && stack[len(stack)-1] == popped[j] {
			stack = stack[:len(stack)-1]
			j++
		}
		i++
	}
	return len(stack) == 0
}
