package lru

/**
146
运用你所掌握的数据结构，设计和实现一个  LRU (最近最少使用) 缓存机制 。 实现 LRUCache 类：

LRUCache(int capacity) 以正整数作为容量 capacity 初始化 LRU 缓存
int get(int key) 如果关键字 key 存在于缓存中，则返回关键字的值，否则返回 -1 。
void put(int key, int value) 如果关键字已经存在，则变更其数据值；
如果关键字不存在，则插入该组「关键字-值」。
当缓存容量达到上限时，它应该在写入新数据之前删除最久未使用的数据值，从而为新的数据值留出空间。

思想：
在 container/list 中，这个双向链表的每个结点的类型是 Element。
Element 中存了 4 个值，前驱和后继结点，双向链表的头结点，value 值。
这里的 value 是 interface 类型。笔者在这个 value 里面存了 pair 这个结构体。
这就解释了 list 里面存的是什么数据。

为什么要存 pair 呢？单单只存 v 不行么，为什么还要存一份 key ？
原因是在 LRUCache 执行删除操作的时候，需要维护 2 个数据结构，一个是 map，一个是双向链表。
在双向链表中删除淘汰出去的 value，在 map 中删除淘汰出去 value 对应的 key。
如果在双向链表的 value 中不存储 key，那么再删除 map 中的 key 的时候有点麻烦。
如果硬要实现，需要先获取到双向链表这个结点 Element 的地址。
然后遍历 map，在 map 中找到存有这个 Element 元素地址对应的 key，再删除。
这样做时间复杂度是 O(n)，做不到 O(1)。所以双向链表中的 Value 需要存储这个 pair。

LRUCache 的 Get 操作很简单，在 map 中直接读取双向链表的结点。
如果 map 中存在，将它移动到双向链表的表头，并返回它的 value 值，如果 map 中不存在，返回 -1。

LRUCache 的 Put 操作也不难。
先查询 map 中是否存在 key，如果存在，更新它的 value，并且把该结点移到双向链表的表头。
如果 map 中不存在，新建这个结点加入到双向链表和 map 中。
最后别忘记还需要维护双向链表的 cap，如果超过 cap，需要淘汰最后一个结点，双向链表中删除这个结点，map 中删掉这个结点对应的 key。

总结，LRU 是由一个 map 和一个双向链表组成的数据结构。
map 中 key 对应的 value 是双向链表的结点。
双向链表中存储 key-value 的 pair。
双向链表表首更新缓存，表尾淘汰缓存。
*/

type LRUCache struct {
	head, tail *Node
	Keys       map[int]*Node
	Cap        int
}

type Node struct {
	Key, Val   int
	Prev, Next *Node
}

func Constructor(capacity int) LRUCache {
	return LRUCache{Keys: make(map[int]*Node), Cap: capacity}
}

func (this *LRUCache) Get(key int) int {
	if node, ok := this.Keys[key]; ok {
		this.Remove(node)
		this.Add(node)
		return node.Val
	}
	return -1
}

func (this *LRUCache) Put(key int, value int) {
	if node, ok := this.Keys[key]; ok {
		node.Val = value
		this.Remove(node)
		this.Add(node)
		return
	} else {
		node = &Node{Key: key, Val: value}
		this.Keys[key] = node
		this.Add(node)
	}
	if len(this.Keys) > this.Cap {
		delete(this.Keys, this.tail.Key)
		this.Remove(this.tail)
	}
}

func (this *LRUCache) Add(node *Node) {
	node.Prev = nil
	node.Next = this.head
	if this.head != nil {
		this.head.Prev = node
	}
	this.head = node
	if this.tail == nil {
		this.tail = node
		this.tail.Next = nil
	}
}

func (this *LRUCache) Remove(node *Node) {
	if node == this.head {
		this.head = node.Next
		node.Next = nil
		return
	}
	if node == this.tail {
		this.tail = node.Prev
		node.Prev.Next = nil
		node.Prev = nil
		return
	}
	node.Prev.Next = node.Next
	node.Next.Prev = node.Prev
}