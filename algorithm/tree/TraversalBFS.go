package tree

import (
	"container/list"
	"fmt"
	"math"
	"sort"
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


/**
515
给定一棵二叉树的根节点 root ，请找出该二叉树中每一层的最大值。
**/
// 解法一 层序遍历二叉树，再将每层排序取出最大值
func largestValues(root *TreeNode) []int {
	tmp := levelOrderBySlice(root)
	res := []int{}
	for i := 0; i < len(tmp); i++ {
		sort.Ints(tmp[i])
		res = append(res, tmp[i][len(tmp[i])-1])
	}
	return res
}

// 解法二 层序遍历二叉树，遍历过程中不断更新最大值
func largestValues1(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}
	q := []*TreeNode{root}
	var res []int
	for len(q) > 0 {
		qlen := len(q)
		max := math.MinInt32
		for i := 0; i < qlen; i++ {
			node := q[0]
			q = q[1:]
			if node.Val > max {
				max = node.Val
			}
			if node.Left != nil {
				q = append(q, node.Left)
			}
			if node.Right != nil {
				q = append(q, node.Right)
			}
		}
		res = append(res, max)
	}
	return res
}

// 解法三 深度遍历二叉树
func largestValues3(root *TreeNode) []int {
	var res []int
	var dfs func(root *TreeNode, level int)
	dfs = func(root *TreeNode, level int) {
		if root == nil {
			return
		}
		if len(res) == level {
			res = append(res, root.Val)
		}
		if res[level] < root.Val {
			res[level] = root.Val
		}

		dfs(root.Right, level+1)
		dfs(root.Left, level+1)
	}
	dfs(root, 0)
	return res
}

