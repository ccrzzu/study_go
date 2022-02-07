package main

import (
	"fmt"

	"github.com/t924417424/BloomFilter"
)

/**
 *布隆过滤器本质上是一个数据结构，它可以用来判断某个元素是否在集合内，
 *具有运行快速，内存占用小的特点。 而高效插入和查询的代价就是，
 *Bloom Filter 是一个基于概率的数据结构：
 *它只能告诉我们一个元素绝对不在集合内或可能在集合内。
**/
func main() {
	filter := bloomfilter.NewBloomfilter(bloomfilter.Config{HashLoop: 30})
	filter.Insert([]byte("123"))
	println(filter.Contains([]byte("123")))
	//filter.Debug()

	myFilter := bloomfilter.NewBloomfilter(bloomfilter.Config{HashLoop: 20})
	key := []byte("123")
	key2 := []byte("456")
	key3 := []byte("999")
	fmt.Println(string(key3))
	myFilter.Insert(key)
	myFilter.Insert(key2)
	println(myFilter.Contains(key))
	println(myFilter.Contains(key2))
	println(myFilter.Contains(key3))
}
