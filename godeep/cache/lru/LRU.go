package lru

import (
	"container/list"
	"fmt"
	"runtime"
)

/**
LRU（least recently used）最近最少使用算法

如上图示，实现 lru算法 的缓存架构图：

lru算法 是相对平衡的一种算法。
核心原则是：如果数据最近被访问过，那么将来被访问的概率会更高
如上图，用双链表来实现的话，如果某条数据被访问了，则把该条数据移动到链表尾部，
队尾是最少使用的元素，内存超出限制时，淘汰队尾元素即可

1. map 用来存储键值对。这是实现缓存最简单直接的数据结构，因为它的查找记录和增加记录时间复杂度都是 O(1)

2. list.List 是go标准库提供的双链表。
通过这个双链表存放具体的值，移动任意记录到队首的时间复杂度都是 O(1)，
在队首增加记录的时间复杂度是 O(1)，删除任意一条记录的时间复杂度是 O(1)
*/

// 定义cache接口
type Cache interface {
	// 设置/添加一个缓存，如果key存在，则用新值覆盖旧值
	Set(key string, value interface{})
	// 通过key获取一个缓存值
	Get(key string) interface{}
	// 通过key删除一个缓存值
	Del(key string)
	// 删除 '最无用' 的一个缓存值
	DelOldest()
	// 获取缓存已存在的元素个数
	Len() int
	// 缓存中 元素 已经所占用内存的大小
	UseBytes() int
}

// 结构体，数组，切片，map,要求实现 Value 接口，该接口只有1个 Len 方法，返回占用内存的字节数
type Value interface {
	Len() int
}

// 定义key,value 结构
type entry struct {
	key   string
	value interface{}
}

// 计算出元素占用内存字节数
func (e *entry) Len() int {
	return CalcLen(e.value)
}

// 计算value占用内存大小
func CalcLen(value interface{}) int {
	var n int
	switch v := value.(type) {
	case Value: // 结构体，数组，切片，map,要求实现 Value 接口，该接口只有1个 Len 方法，返回占用的内存字节数，如果没有实现该接口，则panic
		n = v.Len()
	case string:
		if runtime.GOARCH == "amd64" {
			n = 16 + len(v)
		} else {
			n = 8 + len(v)
		}
	case bool, int8, uint8:
		n = 1
	case int16, uint16:
		n = 2
	case int32, uint32, float32:
		n = 4
	case int64, uint64, float64:
		n = 8
	case int, uint:
		if runtime.GOARCH == "amd64" {
			n = 8
		} else {
			n = 4
		}
	case complex64:
		n = 8
	case complex128:
		n = 16
	default:
		panic(fmt.Sprintf("%T is not implement cache.value", value))
	}

	return n
}

type lru struct {
	// 缓存最大容量，单位字节，这里值最大存放的 元素 的容量，key不算
	maxBytes int

	// 已使用的字节数，只包括value， key不算
	usedBytes int

	// 双链表
	ll *list.List
	// map的key是字符串，value是双链表中对应节点的指针
	cache map[string]*list.Element
}

// TODO: 定义fifo结构体
type fifo struct {
	// 缓存最大容量，单位字节
	// groupCache 使用的是最大存放 entry 个数
	maxBytes int

	// 已使用的字节数，只包括值， key不算
	usedBytes int

	// 双链表
	ll *list.List
	// map的key是字符串，value是双链表中对应节点的指针
	cache map[string]*list.Element
}

// 创建一个新 Cache，如果 maxBytes 是0，则表示没有容量限制
func NewLruCache(maxBytes int) Cache {
	return &fifo{
		maxBytes: maxBytes,
		ll:       list.New(),
		cache:    make(map[string]*list.Element),
	}
}

// 通过 Set 方法往 Cache 头部增加一个元素，如果已经存在，则移到头部，并更新值
func (l *lru) Set(key string, value interface{}) {
	if element, ok := l.cache[key]; ok {
		// 移动到头部
		l.ll.MoveToFront(element)
		eVal := element.Value.(*entry)
		// 重新计算内存占用
		l.usedBytes = l.usedBytes - CalcLen(eVal.value) + CalcLen(value)
		// 更新value
		element.Value = value
	} else {
		element := &entry{
			key:   key,
			value: value,
		}

		e := l.ll.PushFront(element) // 头部插入一个元素并返回该元素
		l.cache[key] = e
		// 计算内存占用
		l.usedBytes += element.Len()
	}

	// 如果超出内存长度，则删除队首的节点. 0表示无内存限制
	for l.maxBytes > 0 && l.maxBytes < l.usedBytes {
		l.DelOldest()
	}
}

// 获取指定元素（有访问要将该元素移动到头部）
func (l *lru) Get(key string) interface{} {
	if e, ok := l.cache[key]; ok {
		// 移动到头部
		l.ll.MoveToFront(e)
		return e.Value.(*entry).value
	}

	return nil
}

// 删除指定元素
func (l *lru) Del(key string) {
	if e, ok := l.cache[key]; ok {
		l.removeElement(e)
	}
}

// 删除最 '无用' 元素，链表尾部为最无用元素
func (l *lru) DelOldest() {
	l.removeElement(l.ll.Back())
}

// 删除元素并更新内存占用大小
func (l *lru) removeElement(e *list.Element) {
	if e == nil {
		return
	}

	l.ll.Remove(e)
	en := e.Value.(*entry)
	l.usedBytes -= en.Len()
	delete(l.cache, en.key)
}

// 缓存池元素数量
func (l *lru) Len() int {
	return l.ll.Len()
}

// 缓存池已经占用的内存大小
func (l *lru) UseBytes() int {
	return l.usedBytes
}

// TODO: 通过 Set 方法往 Cache 头部增加一个元素（如果已经存在，则移到头部，并修改值）
func (f *fifo) Set(key string, value interface{}) {
	if element, ok := f.cache[key]; ok {
		f.ll.MoveToFront(element)
		eVal := element.Value.(*entry)
		f.usedBytes = f.usedBytes - CalcLen(eVal.value) + CalcLen(value) // 更新占用内存大小
		element.Value = value
	} else {
		element := &entry{key, value}
		e := f.ll.PushFront(element) // 头部插入一个元素并返回该元素
		f.cache[key] = e

		f.usedBytes += element.Len()
	}

	// 如果超出内存长度，则删除队首的节点
	for f.maxBytes > 0 && f.maxBytes < f.usedBytes {
		f.DelOldest()
	}
}

// TODO: 获取指定元素
func (f *fifo) Get(key string) interface{} {
	if e, ok := f.cache[key]; ok {
		return e.Value.(*entry).value
	}

	return nil
}

// TODO: 删除指定元素
func (f *fifo) Del(key string) {
	if e, ok := f.cache[key]; ok {
		f.removeElement(e)
	}
}

// TODO: 删除最 '无用' 元素
func (f *fifo) DelOldest() {
	f.removeElement(f.ll.Back())
}

// TODO: 删除元素并更新内存占用大小
func (f *fifo) removeElement(e *list.Element) {
	if e == nil {
		return
	}

	f.ll.Remove(e)
	en := e.Value.(*entry)
	f.usedBytes -= en.Len()
	delete(f.cache, en.key)
}

// TODO: 缓存中元素个数
func (f *fifo) Len() int {
	return f.ll.Len()
}

// TODO: 缓存池占用内存大小
func (f *fifo) UseBytes() int {
	return f.usedBytes
}
