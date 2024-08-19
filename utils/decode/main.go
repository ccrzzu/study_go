package main

import (
	//"encoding/binary"
	"fmt"
	"os"
	"time"
)

func main() {
	// dataLengthByte := []byte{255, 255, 255, 255}
	// fmt.Println(binary.LittleEndian.Uint32(dataLengthByte))

	// return
	decoder := NewDecoder()

	go func() {
		defer recoverCoreDump()
		var i int
		for nsPayload := range decoder.Result() {
			//fmt.Println(nsPayload)
			i++
			fmt.Printf("success-%d: %s\n", i, string(nsPayload))
		}
	}()
	//读取文件
	content, err := os.ReadFile("/Users/duyao/Downloads/b.txt")
	if err != nil {
		fmt.Println(err)
	}
	//fmt.Println(content)
	//
	decoder.Feed(content)
	time.Sleep(time.Second * 30)
	//[]byte := []byte{}
	//dataLength := int(binary.LittleEndian.Uint32(dataLengthByte))
	//fmt.Println()
}

func recoverCoreDump() {
	if err := recover(); err != nil {
		fmt.Println(err)
	}
}
