package tree

import (
	"strconv"
	"strings"
)

var res [][]int

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
