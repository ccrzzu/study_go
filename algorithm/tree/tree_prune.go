package tree


//树的剪枝
func pruneTree(root *TreeNode) *TreeNode {
	return deal(root)
}

//树的剪枝
func deal(node *TreeNode) *TreeNode {
	if node == nil {
		return nil
	}
	node.Left = deal(node.Left)
	node.Right = deal(node.Right)
	//如果这个节点的左右节点都可以剪，且当前节点值也为0，则当前节点为nil，可被整体剪
	if node.Left == nil && node.Right == nil && node.Val == 0 {
		return nil
	}
	return node
}