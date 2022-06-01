package tree

import "math"

/**
124
路径 被定义为一条从树中任意节点出发，沿父节点-子节点连接，达到任意节点的序列。
同一个节点在一条路径序列中 至多出现一次 。该路径 至少包含一个 节点，且不一定经过根节点。

路径和 是路径中各节点值的总和。

给你一个二叉树的根节点 root ，返回其 最大路径和 。
*/

//二叉树的最大路径和
//思想计算每个节点的最大贡献值，就是他本身加上左右子树不为负的最大贡献值
//然后最大路径和就是本身节点的值加上左右子树的最大贡献值之和
//设定一个变量，一直更新即可得到最大路径和
func maxPathSum(root *TreeNode) int {
	sum := math.MinInt64
	var maxGain func(*TreeNode) int
	maxGain = func(node *TreeNode) int {
		if node == nil {
			return 0
		}
		leftGain := max(maxGain(node.Left), 0)
		rightGain := max(maxGain(node.Right), 0)
		sum = max(sum, node.Val+leftGain+rightGain)
		return node.Val + max(leftGain, rightGain)
	}
	maxGain(root)
	return sum
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
