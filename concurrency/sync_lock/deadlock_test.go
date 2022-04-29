package concurrency

import (
	"fmt"
	"math/rand"
	"runtime"
	"sync"
	"testing"
	"time"
)

func TestDeadLock(t *testing.T) {
	ch := make(chan int)
	ch <- 1
	//<-ch
	go func() {
		<-ch
	}()
}

func TestDeadLock2(t *testing.T) {
	ch1 := make(chan int)
	ch2 := make(chan int)
	go func() {
		for {
			select {
			case <-ch1:
				ch2 <- 123
			}

		}
	}()
	for {
		select {
		case <-ch2:
			ch1 <- 123
		}
	}
}

func TestDeadLock3(t *testing.T) {
	lock1 := sync.Mutex{}
	lock2 := sync.Mutex{}
	lock1.Lock()
	lock2.Lock()
	go func() {
		lock2.Lock()
		lock1.Unlock()
	}()
	go func() {
		lock1.Lock()
		lock2.Unlock()
	}()
	for { // 防止 主 go 程 退出
		runtime.GC()
	}
}

// 使用读写锁
var rwMutex2 sync.RWMutex

func readGo2(idx int, in <-chan int) { // 读go程
	for {
		time.Sleep(time.Millisecond * 500) // 放大实验现象// 一个go程可以读 无限 次。
		rwMutex2.RLock()                   // 读模式加  读写锁
		num := <-in                        // 从 公共的 channel 中获取数据
		fmt.Printf("%dth 读 go程，读到：%d\n", idx, num)
		rwMutex2.RUnlock() // 解锁 读写锁
	}
}

func writeGo2(idx int, out chan<- int) {
	for { // 一个go程可以写 无限 次。
		// 生产一个随机数
		num := rand.Intn(500)
		rwMutex2.Lock() // 写模式加  读写锁
		out <- num
		fmt.Printf("-----%dth 写 go程，写入：%d\n", idx, num)
		rwMutex2.Unlock() // 解锁  读写锁

		//time.Sleep(time.Millisecond * 200)        // 放大实验现象
	}
}

func TestDeadLock4(t *testing.T) {
	rand.Seed(time.Now().UnixNano())

	ch := make(chan int, 5)
	for i := 0; i < 5; i++ {
		go readGo2(i+1, ch)
	}

	for i := 0; i < 5; i++ {
		go writeGo2(i+1, ch)
	}

	for {
		runtime.GC()
	}
}
