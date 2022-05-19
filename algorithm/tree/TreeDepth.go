package tree

import (
	"container/list"
	"math"
)

//二叉树深度 BFS解法
func minDepthByBFS(root *TreeNode) int {
	if root == nil {
		return 0
	}
	l := list.New()
	l.PushFront(root)
	depth := 1
	for l.Len() > 0 {
		lLen := l.Len()
		for i := 0; i < lLen; i++ {
			node := l.Remove(l.Back()).(*TreeNode)
			if node.Left == nil && node.Right == nil {
				return depth
			}
			if node.Left != nil {
				l.PushFront(node.Left)
			}
			if node.Right != nil {
				l.PushFront(node.Right)
			}
		}
		depth++
	}
	return depth
}

//二叉树深度 DFS 即递归解法
func minDepthByDFSorDG(root *TreeNode) int {
	if root == nil {
		return 0
	}
	if root.Left == nil && root.Right == nil {
		return 1
	}
	minD := math.MaxInt32
	if root.Left != nil {
		minD = int(math.Min(float64(minDepthByDFSorDG(root.Left)), float64(minD)))
	}
	if root.Right != nil {
		minD = int(math.Min(float64(minDepthByDFSorDG(root.Right)), float64(minD)))
	}
	return minD + 1
}

//求解一棵树的深度
func MaxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return int(math.Max(float64(MaxDepth(root.Left)), float64(MaxDepth(root.Right)))) + 1
}
