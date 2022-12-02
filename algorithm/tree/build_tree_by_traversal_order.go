package tree

//重建二叉树：从前序和中序遍历中构造二叉树
/**
 * 递归逻辑
 * 由于同一颗子树的前序遍历和中序遍历的长度显然是相同的，前序第一个节点肯定是根节点，
 * 然后可以找到中序遍历的根节点值所在索引来得到左右子树的数量，继而得到前序遍历左右子树的索引范围
 * 举例：
 * 前序：1、2、3、4、5
 * 中序：2、1、4、3、5
 * 所以可以根据1所在中序的位置，将中序分成2堆 ：左子树是2和右子树是4、3、5
 * 同时根据上一步分成的左右子树的数量，可得出前序遍历里归属于左右子树的序列：左子树是2和右子树是3、4、5
 * 最后通过递归使用当前函数完成构建
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
/**
 * 中序：2、1、4、3、5
 * 后序：2、4、5、3、1
 * 因为后序遍历的最后一位一定是root的位置
 * 所以可以根据1所在中序的位置，将中序分成2堆 ：左子树是2和右子树是4、3、5
 * 同时根据上一步分成的左右子树的数量，可得出后序遍历里归属于左右子树的序列：左子树是2和右子树是4、5、3
 * 到这里，就可以分别对这两个左右子树递归的使用构建函数完成构建
 */
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
