package main

import (
	"fmt"
	"time"
)

func main() {
	for i := 0; i < 10; i++ {
		printNum(i)
	}
}

func printNum(i int){
	defer fmt.Println(i)
	time.Sleep(1 * time.Second)
}
