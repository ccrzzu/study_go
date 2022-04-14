package consistent_hash

type Node struct {
	Id     string
	hashId uint32
}

func NewNode(realNodeId, virtaulNodeId string) *Node {
	return &Node{
		Id:     realNodeId,
		hashId: buildHashId(virtaulNodeId),
	}
}

type Nodes []*Node

func (n Nodes) Len() int           { return len(n) }
func (n Nodes) Swap(i, j int)      { n[i], n[j] = n[j], n[i] }
func (n Nodes) Less(i, j int) bool { return n[i].hashId < n[j].hashId }