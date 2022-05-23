package tree

//将二叉树展开成链表
func flatten(root *TreeNode) {
	if root == nil {
		return
	}
	flatten(root.Left)
	flatten(root.Right)

	//1.左右子树已经被展开成链表的情况下
	left := root.Left
	right := root.Right
	//将左子树先变成右子树
	root.Left = nil
	root.Right = left

	//2.将root原先的右子树挂到原先左子树的最后节点
	p := root
	for p.Right != nil {
		p = p.Right
	}
	p.Right = right
}