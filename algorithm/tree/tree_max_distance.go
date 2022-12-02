package tree

import "math"

//二叉树中任意2个节点之间最远距离，也就是边的数目最大
func diameterOfBinaryTree(root *TreeNode) int {
	sum := 0
	var dfs func(*TreeNode) int
	dfs = func(node *TreeNode) int {
		if node == nil {
			return 0
		}
		left := dfs(node.Left)
		right := dfs(node.Right)
		sum = int(math.Max(float64(sum), float64(left+right)))
		return int(math.Max(float64(left), float64(right))) + 1
	}
	dfs(root)
	return sum
}
