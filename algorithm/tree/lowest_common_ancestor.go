package tree

//最近公共祖先
/**
236
给定一个二叉树, 找到该树中两个指定节点的最近公共祖先。
百度百科中最近公共祖先的定义为：“对于有根树 T 的两个结点 p、q，最近公共祖先表示为一个结点 x，
满足 x 是 p、q 的祖先且 x 的深度尽可能大（一个节点也可以是它自己的祖先）。”
解题思路 #
这是一套经典的题目，寻找任意一个二叉树中两个结点的 LCA 最近公共祖先，考察递归
*/
func LowestCommonAncestorInBinaryTree(root, p, q *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	if root == p || root == q {
		return root
	}
	left := LowestCommonAncestorInBinaryTree(root.Left, p, q)
	right := LowestCommonAncestorInBinaryTree(root.Right, p, q)
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
func lowestCommonAncestor236(root, p, q *TreeNode) *TreeNode {
	if root == nil || root == q || root == p {
		return root
	}
	left := lowestCommonAncestor236(root.Left, p, q)
	right := lowestCommonAncestor236(root.Right, p, q)
	if left != nil {
		if right != nil {
			return root
		}
		return left
	}
	return right
}

/**
235
给定一个二叉搜索树, 找到该树中两个指定节点的最近公共祖先。
百度百科中最近公共祖先的定义为：“对于有根树 T 的两个结点 p、q，最近公共祖先表示为一个结点 x，
满足 x 是 p、q 的祖先且 x 的深度尽可能大（一个节点也可以是它自己的祖先）。”
解题思路 #
在二叉搜索树中求两个节点的最近公共祖先，由于二叉搜索树的特殊性质，所以找任意两个节点的最近公共祖先非常简单。
*/
func lowestCommonAncestorInBST(root, p, q *TreeNode) *TreeNode {
	if p == nil || q == nil || root == nil {
		return nil
	}
	if p.Val < root.Val && q.Val < root.Val {
		return lowestCommonAncestorInBST(root.Left, p, q)
	}
	if p.Val > root.Val && q.Val > root.Val {
		return lowestCommonAncestorInBST(root.Right, p, q)
	}
	return root
}