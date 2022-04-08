package main

import "runtime"

//连续栈替代了分段栈后，当栈空间不足时，会申请初始化一片新的2倍于旧栈空间的大小的内存区域，
//并将原栈中的所有值都迁移到新栈中。
/**
    使用连续栈机制时，栈空间不足导致的扩容会经历以下几个步骤：
	1、调用用runtime.newstack在内存空间中分配更大的栈内存空间；
	2、使用runtime.copystack将旧栈中的所有内容复制到新的栈中；
	3、将指向旧栈对应变量的指针重新指向新栈；
	4、调用runtime.stackfree销毁并回收旧栈的内存空间；
*/
func main() {
	// 栈扩容
	/* var x [10]int
	println(cap(x), &x)
	a(x)
	println(cap(x), &x) */

	// 栈缩容
	var x [10]int
	println(&x)
	a(x)
	runtime.GC()
	println(&x)
}

//go:noinline
func a(x [10]int) {
	println(`func a`, &x)
	var y [100]int
	b(y)
}

//go:noinline
func b(x [100]int) {
	println(`func b`)
	var y [1000]int
	c(y)
}

//go:noinline
func c(x [1000]int) {
	println(`func c`)
}
