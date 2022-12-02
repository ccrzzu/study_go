package tree


//求解一颗树的节点数目
func countNodes(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return 1 + countNodes(root.Left) + countNodes(root.Right)
}

//求一颗完全二叉树的节点数目
func countNodes2(root *TreeNode) int {
	if root == nil {
		return 0
	}
	leftLevel := countLevel(root.Left)
	rightLevel := countLevel(root.Right)
	if leftLevel == rightLevel {
		return countNodes(root.Right) + (1 << leftLevel)
	} else {
		return countNodes(root.Left) + (1 << rightLevel)
	}
}

//判断左子树的层数
func countLevel(root *TreeNode) int {
	var level int
	for root != nil {
		level++
		root = root.Left
	}
	return level
}