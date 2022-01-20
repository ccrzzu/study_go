package consistent_hash

import (
	"sort"
	"strconv"
	"sync"
)

type Ring struct {
	replicas int
	Nodes    Nodes
	sync.Mutex
}

func NewRing(replicas int) *Ring {
	return &Ring{Nodes: Nodes{}, replicas: replicas}
}

func (r *Ring) AddNodes(ids []string) {
	r.Lock()
	defer r.Unlock()

	r.buildNodeIds(ids, func(readlNodeId, virtualNodeId string) {
		node := NewNode(readlNodeId, virtualNodeId)
		r.Nodes = append(r.Nodes, node)
	})

	sort.Sort(r.Nodes)
}

func (r *Ring) RemoveNodes(ids []string) error {
	r.Lock()
	defer r.Unlock()

	ok := true
	r.buildNodeIds(ids, func(readlNodeId, virtualNodeId string) {
		nodeIndex := r.search(virtualNodeId)
		if nodeIndex >= r.Nodes.Len() || r.Nodes[nodeIndex].Id != readlNodeId {
			ok = false
		} else {
			r.Nodes = append(r.Nodes[:nodeIndex], r.Nodes[nodeIndex+1:]...)
		}
	})

	if !ok {
		return NodeNotFoundErr
	}
	return nil
}

func (r *Ring) GetNode(id string) *Node {
	i := r.search(id)
	if i >= r.Nodes.Len() {
		i = 0
	}

	return r.Nodes[i]
}

func (r *Ring) search(id string) int {
	return sort.Search(r.Nodes.Len(), func(i int) bool {
		return r.Nodes[i].hashId >= buildHashId(id)
	})
}

func (r *Ring) buildNodeIds(ids []string, fn func(realNodeId, virtualNodeId string)) {
	for _, id := range ids {
		for i := 0; i < r.replicas; i++ {
			fn(id, strconv.Itoa(i)+id)
		}
	}
}
