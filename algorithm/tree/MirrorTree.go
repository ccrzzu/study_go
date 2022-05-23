package tree

/**
检查一棵树是不是镜像树
*/
func CheckIsMirrorTree(root *TreeNode) bool {
	return check(root, root)
}

func check(root1, root2 *TreeNode) bool {
	if root1 == nil && root2 == nil {
		return true
	}
	if root1 == nil || root2 == nil {
		return false
	}
	if root1.Val != root2.Val {
		return false
	}
	return check(root1.Left, root1.Right) && check(root1.Right, root2.Left)
}


