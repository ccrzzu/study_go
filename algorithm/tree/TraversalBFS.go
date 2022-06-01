package tree

import (
	"container/list"
	"fmt"
	"study_go/algorithm/array"
)

//BFS: breadth first search  树的广度优先遍历，或者叫按层序遍历

/**
102
给你二叉树的根节点 root ，返回其节点值的 层序遍历 。
（即逐层地，从左到右访问所有节点）。
*/

//解法一，利用slice做队列，bfs广度优先遍历
func levelOrderBySlice(root *TreeNode) [][]int {
	result := make([][]int, 0)
	if root == nil {
		return result
	}
	queue := []*TreeNode{root}
	for len(queue) > 0 {
		currentLevelNode := make([]int, 0)
		queueLength := len(queue)
		for i := 0; i < queueLength; i++ {
			node := queue[0]
			queue = queue[1:]
			currentLevelNode = append(currentLevelNode, node.Val)
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
		result = append(result, currentLevelNode)
	}
	return result
}

//解法二：BFS广度优先遍历解法
func levelOrderByGoList(root *TreeNode) [][]int {
	result := make([][]int, 0)
	if root == nil {
		return result
	}
	queue := list.New()
	queue.PushFront(root)
	for queue.Len() > 0 {
		currentLevelNode := make([]int, 0)
		queueLength := queue.Len()
		for i := 0; i < queueLength; i++ {
			node := queue.Remove(queue.Back()).(*TreeNode)
			currentLevelNode = append(currentLevelNode, node.Val)
			if node.Left != nil {
				queue.PushFront(node.Left)
			}
			if node.Right != nil {
				queue.PushFront(node.Right)
			}
		}
		result = append(result, currentLevelNode)
	}
	return result
}

//解法三：递归解法
func levelOrderByDG(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}
	var res [][]int
	var dfs func(*TreeNode, int)
	dfs = func(node *TreeNode, level int) {
		if node == nil {
			return
		}
		if level == len(res) {
			res = append(res, []int{})
		}
		fmt.Println(res)
		res[level] = append(res[level], node.Val)
		dfs(node.Left, level+1)
		dfs(node.Right, level+1)
	}
	dfs(root, 0)
	return res
}

/**
103 之字形遍历二叉树
给你二叉树的根节点 root ，返回其节点值的 锯齿形层序遍历 。
（即先从左往右，再从右往左进行下一层遍历，以此类推，层与层之间交替进行）。
*/
func levelOrderInZhi(root *TreeNode) [][]int {
	result := [][]int{}
	if root == nil {
		return result
	}
	queue := []*TreeNode{root}
	flag := true
	for len(queue) > 0 {
		currentLevel := make([]int, 0)
		qlen := len(queue)
		for i := 0; i < qlen; i++ {
			node := queue[0]
			queue = queue[1:]
			currentLevel = append(currentLevel, node.Val)
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
		if flag {
			result = append(result, currentLevel)
		} else {
			result = append(result, array.ReverseArray(currentLevel))
		}
		flag = !flag
	}
	return result
}

//第二种写法
func levelOrderInZhi2(root *TreeNode) [][]int {
	res := [][]int{}
	if root == nil {
		return res
	}
	queue := []*TreeNode{root}
	for level := 0; len(queue) > 0; level++ {
		vals := []int{}
		lenQueue := len(queue)
		for i := 0; i < lenQueue; i++ {
			node := queue[0]
			queue = queue[1:]
			vals = append(vals, node.Val)
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
		if level%2 == 1 {
			for i := 0; i < len(vals)/2; i++ {
				vals[i], vals[len(vals)-i-1] = vals[len(vals)-i-1], vals[i]
			}
		}
		res = append(res, vals)
	}
	return res
}
