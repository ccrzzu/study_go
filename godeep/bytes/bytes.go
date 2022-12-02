package main

import "bytes"

func main(){
	// 字符串转换为字节数组
	buff3 := []byte("hello world hello world")
	seq := []byte("hello")
	
	// Count counts the number of non-overlapping instances of sep in s
	res := bytes.Count(buff3, seq)
	println(res)

	// Contains reports whether subslice is within b
	contains := bytes.Contains(buff3, seq) //true
	println(contains)

	res = bytes.Index(buff3, seq) // 0
	println(res)

	res = bytes.LastIndex(buff3, seq)
	println(res)

	/**
	Rune literals are just an integer value (as you've written).
	They are "mapped" to their unicode codepoint.
	 */
	a := rune('e')
	res = bytes.IndexRune(buff3, a) // 1
	println(res)

	println("------------")
}