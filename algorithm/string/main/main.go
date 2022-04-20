package main

import (
	"fmt"
	"reflect"
	myString "study_go/algorithm/string"
)

func main() {
	str1 := []byte("abcd")
	fmt.Println(string(str1[0]))
	myString.ReverseString(str1)
	fmt.Println(str1)

	str2 := "cdaeb"
	// fmt.Println(str1[0])
	// fmt.Println(str2[0])
	// fmt.Println(str1[0] ^ str2[1])
	res := str2[len(str2)-1]
	fmt.Println("start:",res, reflect.TypeOf(res))
	for i:=0;i<len(str1);i++{
		fmt.Println(reflect.TypeOf(str1[i]))
		res ^= str1[i]
		fmt.Println(res)
		res ^= str2[i]
		fmt.Println(res)
	}
	for _, item := range str1 {
		fmt.Println(item,reflect.TypeOf(item))
	}
	fmt.Println(len("我爱你"))
}
