package tree

import (
	"container/list"
	"fmt"
	"study_go/algorithm/array"
)

//树的广度优先遍历，或者叫按层序遍历

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

//判断一棵树是否为完全二叉树
//完全二叉树的定义就是每一层从左往右都是连续的，不能有空的跳过去
func isCompleteTree(root *TreeNode) bool {
	queue := []*TreeNode{root}
	lastIsNil := false
	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]
		if node == nil {
			lastIsNil = true
		} else {
			if lastIsNil {
				return false
			}
			queue = append(queue, node.Left)
			queue = append(queue, node.Right)
		}
	}
	return true
}

//打开转盘锁第一种方法
//你有一个带有四个圆形拨轮的转盘锁。每个拨轮都有10个数字： '0', '1', '2', '3', '4', '5', '6', '7', '8', '9' 。每个拨轮可以自由旋转：例如把 '9' 变为  '0'，'0' 变为 '9' 。每次旋转都只能旋转一个拨轮的一位数字。
//锁的初始数字为 '0000' ，一个代表四个拨轮的数字的字符串。
//列表 deadends 包含了一组死亡数字，一旦拨轮的数字和列表里的任何一个元素相同，这个锁将会被永久锁定，无法再被旋转。
//字符串 target 代表可以解锁的数字，你需要给出最小的旋转次数，如果无论如何不能解锁，返回 -1。
func OpenLock(deadends []string, target string) int {
	//死亡数字从数组转成map，判断时加快速度
	deadendsMap := map[string]bool{}
	for _, item := range deadends {
		deadendsMap[item] = true
	}
	//记录已经穷举过的密码
	visited := map[string]bool{}
	//var queue []string
	step := 0
	//queue := list.New()
	//queue.PushFront("0000")
	visited["0000"] = true
	queue := []string{"0000"}
	for len(queue) > 0 {
		//qSize := queue.Len()
		qSize := len(queue)
		for i := 0; i < qSize; i++ {
			//cur := queue.Remove(queue.Back()).(string)
			cur := queue[0]
			queue = queue[1:]

			if _, ok := deadendsMap[cur]; ok {
				continue
			}
			if cur == target {
				return step
			}

			for j := 0; j < 4; j++ {
				one := plusOne(cur, j)
				if _, ok := visited[one]; !ok {
					//queue.PushFront(one)
					queue = append(queue, one)
					visited[one] = true
				}
				one = minusOne(cur, j)
				if _, ok := visited[one]; !ok {
					//queue.PushFront(one)
					queue = append(queue, one)
					visited[one] = true
				}
			}
		}
		step++
	}
	return -1
}

// 将 s[j] 向上拨动一次
func plusOne(s string, j int) string {
	bytes := []byte(s)
	if bytes[j] == '9' {
		bytes[j] = '0'
	} else {
		bytes[j] += 1
	}
	return string(bytes)
}

// 将 s[i] 向下拨动一次
func minusOne(s string, j int) string {
	bytes := []byte(s)
	if bytes[j] == '0' {
		bytes[j] = '9'
	} else {
		bytes[j] -= 1
	}
	return string(bytes)
}

//打开转盘锁第二种方法，使用双端队列
func OpenLockByBidirectional(deadends []string, target string) int {
	//死亡数字从数组转成map，判断时加快速度
	deadendsMap := map[string]bool{}
	for _, item := range deadends {
		deadendsMap[item] = true
	}
	//记录已经穷举过的密码
	visited := map[string]bool{}
	visited["0000"] = true
	step := 0
	queueFrom := map[string]bool{"0000": true}
	queueTarget := map[string]bool{target: true}

	for len(queueFrom) > 0 && len(queueTarget) > 0 {
		tmp := map[string]bool{}
		for cur := range queueFrom {
			if _, ok := deadendsMap[cur]; ok {
				continue
			}
			if _, ok := queueTarget[cur]; ok {
				return step
			}
			visited[cur] = true

			for j := 0; j < 4; j++ {
				one := plusOne(cur, j)
				if _, ok := visited[one]; !ok {
					tmp[one] = true
				}
				one = minusOne(cur, j)
				if _, ok := visited[one]; !ok {
					tmp[one] = true
				}
			}
		}
		step++
		queueFrom = queueTarget
		queueTarget = tmp
	}
	return -1
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
