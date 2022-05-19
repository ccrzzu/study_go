package tree

//镜像反转二叉树
func invertTree(root *TreeNode) *TreeNode {
	//base case
	if root == nil {
		return nil
	}

	tmp := root.Left
	root.Left = root.Right
	root.Right = tmp

	//递归调换左右子树
	invertTree(root.Left)
	invertTree(root.Right)

	return root
}