package tree

import (
	"container/list"
	"fmt"
)

//BFS: breadth first search  树的广度优先遍历，或者叫按层序遍历

//层序遍历获取树节点
//BFS广度优先遍历解法
func levelOrder(root *TreeNode) [][]int {
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

//bfs广度优先遍历
func levelOrderByQueue(root *TreeNode) [][]int {
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

//之字形遍历二叉树
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
