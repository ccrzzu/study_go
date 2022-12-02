package tree

//给定一个二叉树，原地将它展开为链表。
/**
解题思路：
要求把二叉树“打平”，按照先根遍历的顺序，把树的结点都放在右结点中。

按照递归和非递归思路实现即可。

递归的思路可以这么想：倒序遍历一颗树，即是先遍历右孩子，然后遍历左孩子，最后再遍历根节点。
*/
func flattenSelf(root *TreeNode) {
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

// 解法一 非递归
func flatten(root *TreeNode) {
	list, cur := []int{}, &TreeNode{}
	//前序遍历
	preorder(root, &list)
	cur = root
	for i := 1; i < len(list); i++ {
		cur.Left = nil
		cur.Right = &TreeNode{Val: list[i], Left: nil, Right: nil}
		cur = cur.Right
	}
	return
}

// 解法二 递归
func flatten1(root *TreeNode) {
	if root == nil || (root.Left == nil && root.Right == nil) {
		return
	}
	flatten(root.Left)
	flatten(root.Right)
	currRight := root.Right
	root.Right = root.Left
	root.Left = nil
	for root.Right != nil {
		root = root.Right
	}
	root.Right = currRight
}

// 解法三 递归
func flatten2(root *TreeNode) {
	if root == nil {
		return
	}
	flatten(root.Right)
	if root.Left == nil {
		return
	}
	flatten(root.Left)
	p := root.Left
	for p.Right != nil {
		p = p.Right
	}
	p.Right = root.Right
	root.Right = root.Left
	root.Left = nil
}
