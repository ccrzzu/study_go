package tree

import (
	"container/list"
	"fmt"
	"math"
	"strconv"
	"strings"
)

var res [][]int

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
type Node struct {
	Val   int
	Left  *Node
	Right *Node
	Next  *Node
}

type Codec struct {
	l []string
}

//树的序列化与反序列化
func Constructor() Codec {
	return Codec{}
}
func serializeDG(root *TreeNode, str string) string {
	if root == nil {
		str += "null,"
	} else {
		str += strconv.Itoa(root.Val) + ","
		str = serializeDG(root.Left, str)
		str = serializeDG(root.Right, str)
	}
	return str
}

func (this *Codec) serialize(root *TreeNode) string {
	return serializeDG(root, "")
}

func (this *Codec) deserializeDG() *TreeNode {
	if len(this.l) == 0 {
		return nil
	}
	if this.l[0] == "null" {
		this.l = this.l[1:]
		return nil
	}
	atoi, _ := strconv.Atoi(this.l[0])
	root := &TreeNode{
		Val: atoi,
	}
	this.l = this.l[1:]
	root.Left = this.deserializeDG()
	root.Right = this.deserializeDG()
	return root
}

func (this *Codec) deserialize(data string) *TreeNode {
	l := strings.Split(data, ",")
	for i := 0; i < len(l); i++ {
		if l[i] != "" {
			this.l = append(this.l, l[i])
		}
	}
	return this.deserializeDG()
}

//验证是否为山峰数组
func ValidMountainArray(A []int) bool {
	i := 0
	for i+1 < len(A) && A[i] < A[i+1] {
		i++
	}
	if i == 0 || i == len(A)-1 {
		return false
	}
	for i+1 < len(A) && A[i] > A[i+1] {
		i++
	}

	return i == len(A)-1
}

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

//将满二叉树水平层相邻节点指向下一个
func connect(root *Node) *Node {
	if root == nil {
		return nil
	}
	connectTwoNode(root.Left, root.Right)
	return root
}

func connectTwoNode(node1, node2 *Node) {
	if node1 == nil || node2 == nil {
		return
	}
	// 将传入的两个节点连接
	node1.Next = node2
	// 连接相同父节点的两个子节点
	connectTwoNode(node1.Left, node1.Right)
	connectTwoNode(node2.Left, node2.Right)

	//连接不同父节点的两个子节点
	connectTwoNode(node1.Right, node2.Left)
}

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

//层序遍历获取树节点
//BFS广度优先遍历解法
func levelOrder(root *TreeNode) [][]int {
	result := make([][]int, 0)
	if root == nil {
		return result
	}
	queue := list.New()
	queue.PushFront(root)
	for queue.Len() > 0 {
		currentLevelNode := make([]int, 0)
		queueLength := queue.Len()
		for i := 0; i < queueLength; i++ {
			node := queue.Remove(queue.Back()).(*TreeNode)
			currentLevelNode = append(currentLevelNode, node.Val)
			if node.Left != nil {
				queue.PushFront(node.Left)
			}
			if node.Right != nil {
				queue.PushFront(node.Right)
			}
		}
		result = append(result, currentLevelNode)
	}
	return result
}

//bfs广度优先遍历
func levelOrderByQueue(root *TreeNode) [][]int {
	result := make([][]int, 0)
	if root == nil {
		return result
	}
	queue := []*TreeNode{root}
	for len(queue) > 0 {
		currentLevelNode := make([]int, 0)
		queueLength := len(queue)
		for i := 0; i < queueLength; i++ {
			node := queue[0]
			queue = queue[1:]
			currentLevelNode = append(currentLevelNode, node.Val)
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
		result = append(result, currentLevelNode)
	}
	return result
}

func levelOrderByDG(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}
	res = [][]int{}
	levelOrderDFS(root, 0)
	return res
}

func levelOrderDFS(root *TreeNode, level int) {
	if root == nil {
		return
	}
	if level == len(res) {
		res = append(res, []int{})
	}
	res[level] = append(res[level], root.Val)
	levelOrderDFS(root.Left, level+1)
	levelOrderDFS(root.Right, level+1)
}

//合并树
func mergeTrees(t1 *TreeNode, t2 *TreeNode) *TreeNode {
	if t1 == nil {
		return t2
	}
	if t2 == nil {
		return t1
	}
	t1.Val += t2.Val
	t1.Left = mergeTrees(t1.Left, t2.Left)
	t1.Right = mergeTrees(t1.Right, t2.Right)
	return t1
}

func mergeTreesByQueue(t1 *TreeNode, t2 *TreeNode) *TreeNode {
	if t1 == nil {
		return t2
	}
	if t2 == nil {
		return t1
	}
	rootNode := &TreeNode{Val: t1.Val + t2.Val}
	queue := []*TreeNode{rootNode}
	queue1 := []*TreeNode{t1}
	queue2 := []*TreeNode{t2}
	for len(queue1) > 0 && len(queue2) > 0 {
		node := queue[0]
		queue = queue[1:]
		node1 := queue1[0]
		queue1 = queue1[1:]
		node2 := queue2[0]
		queue2 = queue2[1:]
		left1, left2 := node1.Left, node2.Left
		right1, right2 := node1.Right, node2.Right
		if left1 != nil || left2 != nil {
			if left1 != nil && left2 != nil {
				left := &TreeNode{Val: left1.Val + left2.Val}
				node.Left = left
				queue = append(queue, left)
				queue1 = append(queue1, left1)
				queue2 = append(queue2, left2)
			} else if left1 != nil {
				node.Left = left1
			} else {
				node.Left = left2
			}
		}
		if right1 != nil || right2 != nil {
			if right1 != nil && right2 != nil {
				right := &TreeNode{Val: right1.Val + right2.Val}
				node.Right = right
				queue = append(queue, right)
				queue1 = append(queue1, right1)
				queue2 = append(queue2, right2)
			} else if right1 != nil {
				node.Right = right1
			} else {
				node.Right = right2
			}
		}
	}
	return rootNode
}

/**
给定一个不含重复元素的整数数组。一个以此数组构建的最大二叉树定义如下：
二叉树的根是数组中的最大元素。
左子树是通过数组中最大值左边部分构造出的最大二叉树。
右子树是通过数组中最大值右边部分构造出的最大二叉树。
通过给定的数组构建最大二叉树，并且输出这个树的根节点。
**/
func constructMaximumBinaryTree(nums []int) *TreeNode {
	if len(nums) == 0 {
		return nil
	}
	maxVal := math.MinInt64
	var index int
	for i := 0; i < len(nums); i++ {
		if nums[i] > maxVal {
			maxVal = nums[i]
			index = i
		}
	}
	root := &TreeNode{
		Val: maxVal,
	}
	root.Left = constructMaximumBinaryTree(nums[0:index])
	root.Right = constructMaximumBinaryTree(nums[index+1:])
	return root
}

//从前序和中序遍历中构造二叉树
/**
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

//利用中序和后序遍历数组来构建一棵二叉树
func buildTree2(inorder []int, postorder []int) *TreeNode {
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
	root.Left = buildTree2(inorder[:rootInorderIndex], postorder[0:leftSize])
	root.Right = buildTree2(inorder[rootInorderIndex+1:], postorder[leftSize:len(postorder)-1])
	return root
}

//寻找重复的子树
var subTreesMap map[string]int
var dRes []*TreeNode

func findDuplicateSubtrees(root *TreeNode) []*TreeNode {
	subTreesMap = map[string]int{}
	dRes = make([]*TreeNode, 0)
	findDuplicateSubtreesDG(root)
	return dRes
}

//寻找重复的子树 递归解法
func findDuplicateSubtreesDG(root *TreeNode) string {
	if root == nil {
		return "#"
	}
	left := findDuplicateSubtreesDG(root.Left)
	right := findDuplicateSubtreesDG(root.Right)
	subTreeStr := left + "," + right + "," + strconv.Itoa(root.Val)
	subTreesMap[subTreeStr]++
	if subTreesMap[subTreeStr] == 2 {
		dRes = append(dRes, root)
	}
	fmt.Println(subTreesMap, dRes)
	return subTreeStr
}

//二叉树深度 BFS解法
func minDepthByBFS(root *TreeNode) int {
	if root == nil {
		return 0
	}
	l := list.New()
	l.PushFront(root)
	depth := 1
	for l.Len() > 0 {
		lLen := l.Len()
		for i := 0; i < lLen; i++ {
			node := l.Remove(l.Back()).(*TreeNode)
			if node.Left == nil && node.Right == nil {
				return depth
			}
			if node.Left != nil {
				l.PushFront(node.Left)
			}
			if node.Right != nil {
				l.PushFront(node.Right)
			}
		}
		depth++
	}
	return depth
}

//二叉树深度 DFS 即递归解法
func minDepthByDFSorDG(root *TreeNode) int {
	if root == nil {
		return 0
	}
	if root.Left == nil && root.Right == nil {
		return 1
	}
	minD := math.MaxInt32
	if root.Left != nil {
		minD = int(math.Min(float64(minDepthByDFSorDG(root.Left)), float64(minD)))
	}
	if root.Right != nil {
		minD = int(math.Min(float64(minDepthByDFSorDG(root.Right)), float64(minD)))
	}
	return minD + 1
}
