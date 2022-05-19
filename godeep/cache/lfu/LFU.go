package lfu

import (
	"container/heap"
	"fmt"
	"runtime"
)

/**
LFU（least frequently used）最少使用算法

LFU（最少使用）算法会淘汰缓存中访问次数最少的数据。
其核心原则是，如果数据过去被访问多次，那么将来被访问的频率也会更高。
在 LFU 实现中，需要维护一个按照访问次数排序的队列，每次访问时，次数加1,队伍重新排序，淘汰时选择访问次数最少的即可

如上图示
'queue'是一个二叉堆实现的队列
'weight'为访问次数

相对FIFO算法，LFU 使用了 堆queue，而不是 双链表。
二叉堆 插入记录，更新记录，删除记录，时间复杂度都是 'O(logN)'
map 用来存储键值对，每次访问键值时，都必须更新weight，因此获取记录时间复杂度也是 'O(logN)'
*/

// 最小堆实现的队列
type queue []*entry

// 队列长度
func (q queue) Len() int {
    return len(q)
}

// '<' 是最小堆，'>' 是最大堆
func (q queue) Less(i, j int) bool {
    return q[i].weight < q[j].weight
}

// 交换元素
func (q queue) Swap(i, j int) {
    // 交换元素
    q[i], q[j] = q[j], q[i]
    // 索引不用交换
    q[i].index = i
    q[j].index = j
}

// append ，*q = oldQue[:n-1] 会导致频繁的内存拷贝
// 实际上，如果使用 LFU算法，处于性能考虑，可以将最大内存限制修改为最大记录数限制
// 这样提前分配好 queue 的容量，再使用交换索引和限制索引的方式来实现 Pop 方法，可以免去频繁的内存拷贝，极大提高性能
func (q *queue) Push(v interface{}) {
    n := q.Len()
    en := v.(*entry)
    en.index = n
    *q = append(*q, en) // 这里会重新分配内存，并拷贝数据
}

func (q *queue) Pop() interface{} {
    oldQue := *q
    n := len(oldQue)
    en := oldQue[n-1]
    oldQue[n-1] = nil // 将不再使用的对象置为nil，加快垃圾回收，避免内存泄漏
    *q = oldQue[:n-1] // 这里会重新分配内存，并拷贝数据
    return en
}

// weight更新后，要重新排序，时间复杂度为 O(logN)
func (q *queue) update(en *entry, val interface{}, weight int) {
    en.value = val
    en.weight = weight
    (*q)[en.index] = en
    // 重新排序
    // 分析思路是把 堆(大D) 的树状图画出来，看成一个一个小的堆(小D)，看改变其中一个值，对 大D 有什么影响
    // 可以得出结论，下沉操作和上沉操作分别执行一次能将 queue 排列为堆
    heap.Fix(q, en.index)
}

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

// 定义元素
type entry struct {
    key    string
    value  interface{}
    weight int // 访问次数
    index  int // queue索引
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

// 定义lfu cache 结构体
type lfu struct {
    // 缓存最大容量，单位字节，这里值最大存放的 元素 的容量，key不算
    maxBytes int

    // 已使用的字节数，只包括value， key不算
    usedBytes int

    // 最小堆实现的队列
    queue *queue
    // map的key是字符串，value是entry
    cache map[string]*entry
}

// 创建一个新 Cache，如果 maxBytes 是0，则表示没有容量限制
func NewLfuCache(maxBytes int) Cache {
    queue := make(queue, 0)
    return &lfu{
        maxBytes: maxBytes,
        queue:    &queue,
        cache:    make(map[string]*entry),
    }
}

// 通过 Set 方法往 Cache 头部增加一个元素，如果存在则更新值
func (l *lfu) Set(key string, value interface{}) {
    if en, ok := l.cache[key]; ok {
        l.usedBytes = l.usedBytes - en.Len() + CalcLen(value) // 更新占用内存长度
        l.queue.update(en, value, en.weight+1)
    } else {
        en := &entry{
            key:   key,
            value: value,
        }

        heap.Push(l.queue, en)  // 插入queue 并重新排序为堆
        l.cache[key] = en       // 插入 map
        l.usedBytes += en.Len() // 更新内存占用

        // 如果超出内存长度，则删除最 '无用' 的元素，0表示无内存限制
        for l.maxBytes > 0 && l.usedBytes >= l.maxBytes {
            l.DelOldest()
        }
    }
}

// 获取指定元素,访问次数加1
func (l *lfu) Get(key string) interface{} {
    if en, ok := l.cache[key]; ok {
        l.queue.update(en, en.value, en.weight+1)
        return en.value
    }
    return nil
}

// 删除指定元素（删除queue和map中的val）
func (l *lfu) Del(key string) {
    if en, ok := l.cache[key]; ok {
        heap.Remove(l.queue, en.index)
        l.removeElement(en)
    }
}

// 删除最 '无用' 元素（删除queue和map中的val）
func (l *lfu) DelOldest() {
    if l.Len() == 0 {
        return
    }
    val := heap.Pop(l.queue)
    l.removeElement(val)
}

// 删除元素并更新内存占用大小
func (l *lfu) removeElement(v interface{}) {
    if v == nil {
        return
    }

    en := v.(*entry)

    delete(l.cache, en.key)
    l.usedBytes -= en.Len()
}

// 缓存池元素个数
func (l *lfu) Len() int {
    return l.queue.Len()
}

// 缓存池占用内存大小
func (l *lfu) UseBytes() int {
    return l.usedBytes
}
