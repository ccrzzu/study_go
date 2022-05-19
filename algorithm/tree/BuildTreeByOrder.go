package tree

//重建二叉树：从前序和中序遍历中构造二叉树
/**
 * 递归逻辑
 * 由于同一颗子树的前序遍历和中序遍历的长度显然是相同的，
 * 所以可以根据中序遍历的根节点所在索引来得到左右子树的数量，继而得到前序遍历左右子树的索引范围
 */
func BuildTree(preorder []int, inorder []int) *TreeNode {
	if len(preorder) == 0 {
		return nil
	}
	root := &TreeNode{Val: preorder[0]}
	rootInorderIndex := 0
	for ; rootInorderIndex < len(inorder); rootInorderIndex++ {
		if inorder[rootInorderIndex] == root.Val {
			break
		}
	}
	leftSize := len(inorder[:rootInorderIndex])
	root.Left = BuildTree(preorder[1:leftSize+1], inorder[:rootInorderIndex])
	root.Right = BuildTree(preorder[leftSize+1:], inorder[rootInorderIndex+1:])
	return root
}

//重建二叉树：利用中序和后序遍历数组来构建一棵二叉树
func BuildTree2(inorder []int, postorder []int) *TreeNode {
	if len(postorder) == 0 {
		return nil
	}
	root := &TreeNode{Val: postorder[len(postorder)-1]}
	rootInorderIndex := 0
	for ; rootInorderIndex < len(inorder); rootInorderIndex++ {
		if inorder[rootInorderIndex] == root.Val {
			break
		}
	}
	leftSize := len(inorder[:rootInorderIndex])
	root.Left = BuildTree2(inorder[:rootInorderIndex], postorder[0:leftSize])
	root.Right = BuildTree2(inorder[rootInorderIndex+1:], postorder[leftSize:len(postorder)-1])
	return root
}
