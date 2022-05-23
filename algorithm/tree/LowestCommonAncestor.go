package tree

//最近公共祖先
func LowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	if root == p || root == q {
		return root
	}
	left := LowestCommonAncestor(root.Left, p, q)
	right := LowestCommonAncestor(root.Right, p, q)
	//分别在左右子树中
	if left != nil && right != nil {
		return root
	}
	//左右子树中都没有
	if left == nil && right == nil {
		return nil
	}
	//左右子树中只一个有
	if left != nil {
		return left
	} else {
		return right
	}
}
