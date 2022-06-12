package main

import (
	"fmt"
)

/**
什么是闭包？
闭包就是能够读取其他函数内部变量的函数。
闭包可以简单理解成"定义在一个函数内部的函数"。
在本质上，闭包是将函数内部和函数外部连接起来的桥梁。

defer应用
在defer函数定义时，对外部变量的引用是有两种方式的，分别是作为函数参数和作为闭包引用。
作为函数参数，则在defer定义时就把值传递给defer，并被cache起来；
作为闭包引用的话，则会在defer函数真正调用时根据整个上下文确定当前的值。
defer后面的语句在执行的时候，函数调用的参数会被保存起来，也就是复制了一份。
真正执行的时候，实际上用到的是这个复制的变量，因此如果此变量是一个“值”，那么就和定义的时候是一致的。
如果此变量是一个“引用”，那么就可能和定义的时候不一致。
*/

func main() {
	i1 := Increase()
	i2 := Increase()
	fmt.Println(i1())
	fmt.Println(i2())

	fmt.Println("+++++++++++++++++++++++")

	fmt.Println(func1())
	fmt.Println(func2())
	fmt.Println(func3())
	fmt.Println(func4())
	fmt.Println(func9())
}

func Increase() func() int {
	n := 0
	return func() int {
		n++
		return n
	}
}

func func1() (i int) {
	i = 1
	defer func() {
		i++
	}()
	i++
	return
}

func func2() int {
	i := 1
	defer func() {
		i++
	}()
	i++
	return i
}
func func3() int {
	i := 1
	defer func(i int) {
		i++
	}(i)
	i++
	return i
}

func func4() (i int) {
	i = 1
	defer func(i int) {
		i++
	}(i)
	i++
	return
}

func func9() (i int) {
	i = 1
	defer fmt.Println("result =>", i, func() int { return i * 2 }())
	i++
	return
}
