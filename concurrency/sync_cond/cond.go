package main

import (
	"fmt"
	"strconv"
	"sync"
	"time"
)

func main() {
	locker := new(sync.Mutex)
	condLocal := sync.NewCond(locker)
	done := false

	condLocal.L.Lock()

	go func() {
		time.Sleep(2e9)
		done = true
		condLocal.Signal() // 2当从goroutine发出信号之后，主goroutine才会继续往下面走
	}()

	if !done {
		condLocal.Wait() // 1主goroutine进入cond.Wait的时候，就会进入等待
	}

	fmt.Println("now done is", done)

	//**************************************************//
	//以下是cond.Broadcast的用法
	for i := 0; i < 40; i++ {
		go test(i)
	}
	time.Sleep(time.Second * 5)
	fmt.Println("after sleep 5 second, start broadcast, wake all......")
	cond.Broadcast() //  下发广播给所有等待的goroutine
	time.Sleep(time.Second * 60)
}

var locker = new(sync.Mutex)
var cond = sync.NewCond(locker)

func test(x int) {
	cond.L.Lock() // 获取锁, 此处为什么所有协程都能获取到锁？
	fmt.Println(strconv.Itoa(x), "block, wait...")
	cond.Wait() // 等待通知  暂时阻塞
	fmt.Println(strconv.Itoa(x), "be waked")
	time.Sleep(time.Second * 1)
	cond.L.Unlock() // 释放锁，不释放的话将只会有一次输出
}
