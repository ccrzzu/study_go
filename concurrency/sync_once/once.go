package main

import (
	"fmt"
	"sync"
	"time"
)

//我们多个goroutine都要过一个操作，但是这个操作我只希望被执行一次，这个时候Once就上场了
func main() {
	var once sync.Once
	onceBody := func() {
		fmt.Println("Only once")
	}
	for i := 0; i < 10; i++ {
		go func() {
			once.Do(onceBody)
		}()
	}
	time.Sleep(3e9)
}
