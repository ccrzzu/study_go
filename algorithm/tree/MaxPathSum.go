package tree

import "math"

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
		leftGain := math.Max(float64(maxGain(node.Left)), 0)
		rightGain := math.Max(float64(maxGain(node.Right)), 0)
		sum = int(math.Max(float64(sum), float64(node.Val)+leftGain+rightGain))
		return node.Val + int(math.Max(leftGain, rightGain))
	}
	maxGain(root)
	return sum
}