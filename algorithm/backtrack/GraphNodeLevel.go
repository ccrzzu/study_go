package backtrack

import "fmt"

func main() {
	// var a int
	// fmt.Scan(&a)
	fmt.Printf("%s", "hello world")
}

//字节面试题，一个节点之间是有层次的
//给出所有节点的所在层次level值
//example: 上下关系
//a b c d
//  d e
//  f g
type Node struct {
	name  string
	nexts []Node
	pres  []Node
	level int
}

func Level(nodes []Node) {
	var dfs func(curNode Node, level int, pres []Node)
	dfs = func(curNode Node, level int, pres []Node) {
		if len(pres) == 0 {
			if level > curNode.level {
				curNode.level = level
			}
			return
		}
		level++
		for _, node := range pres {
			dfs(curNode, level, node.pres)
		}
	}
	for _, oriNode := range nodes {
		level := 1
		dfs(oriNode, level, oriNode.pres)
	}
}