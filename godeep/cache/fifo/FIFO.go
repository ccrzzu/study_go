package fifo

import (
	"container/list"
	"fmt"
	"runtime"
)

/**
FIFO（first in first ）先进先出算法
如上图示，实现 fifo算法 的缓存架构图：

fifo 算法是淘汰缓存中最早添加的记录,即一个数据最先进入缓存，那么也应该最先被删除掉。
算法的实现比较简单：创建一个队列（一般通过双链表实现），新增记录添加到队首，淘汰队尾记录

1. map 用来存储键值对。这是实现缓存最简单直接的数据结构，因为它的查找记录和增加记录时间复杂度都是 O(1)

2. list.List 是go标准库提供的双链表。
通过这个双链表存放具体的值，移动任意记录到队首的时间复杂度都是 O(1)，
在队首增加记录的时间复杂度是 O(1)，删除任意一条记录的时间复杂度是 O(1)
*/

// TODO: 定义cache接口
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

// TODO: 结构体，数组，切片，map,要求实现 Value 接口，该接口只有1个 Len 方法，返回占用内存的字节数
type Value interface {
    Len() int
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

// TODO: 定义key,value 结构
type entry struct {
    key   string
    value interface{}
}

// TODO: 计算出元素占用内存字节数
func (e *entry) Len() int {
    return CalcLen(e.value)
}

// TODO: 计算value占用内存大小
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

// TODO: 构造函数，创建一个新 Cache，如果 maxBytes 是0，则表示没有容量限制
func NewFifoCache(maxBytes int) Cache {
    return &fifo{
        maxBytes: maxBytes,
        ll:       list.New(),
        cache:    make(map[string]*list.Element),
    }
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
