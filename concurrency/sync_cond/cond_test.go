package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"sync"
	"testing"
	"time"
)

/**
还有个sync.Cond是用来控制某个条件下，
goroutine进入等待时期，等待信号到来，然后重新启动。
*/

func TestCond(t *testing.T) {
	lock := &sync.Mutex{}
	cond := sync.NewCond(lock)
	for i := 0; i < 10; i++ {
		runGoroutine(cond, i)
	}

	time.Sleep(1 * time.Millisecond)
	fmt.Println("----------------------------: signal 唤醒单个")
	cond.Signal()

	time.Sleep(1 * time.Millisecond)
	fmt.Println("----------------------------: broadcast 唤醒全部")
	cond.Broadcast()

	time.Sleep(2 * time.Second)
}

func runGoroutine(cond *sync.Cond, i int) {
	go func(cond *sync.Cond, i int) {
		cond.L.Lock()
		for condition() {
			fmt.Println("-goroutine-" + strconv.Itoa(i) + " 命中wait")
			cond.Wait()
		}
		fmt.Println("-goroutine-" + strconv.Itoa(i) + " 命中条件")
		cond.L.Unlock()
	}(cond, i)
}

func condition() bool {
	rand.Intn(50)
	if rand.Intn(50) > 20 {
		fmt.Print(true)
		return true
	}
	fmt.Print(false)
	return false
}
