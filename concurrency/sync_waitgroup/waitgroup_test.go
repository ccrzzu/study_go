package concurrency

import (
	"fmt"
	"sync"
	"testing"
)

/**
一个goroutine需要等待一批goroutine执行完毕以后才继续执行，
那么这种多线程等待的问题就可以使用WaitGroup了。
*/

type UserInfo struct {
	UserID     int
	UserName   string
	UserRemain int
}

func TestWaitGroup(t *testing.T) {
	var wg sync.WaitGroup
	userInfoList := []*UserInfo{}
	//wg.Add(len(userIdList)) //另一种方式
	userIdList := []int{1, 2, 3}
	for _, userId := range userIdList {
		wg.Add(1)
		go func(id int, w *sync.WaitGroup) {
			//defer wg.Done() //另一种方式
			userInfoList = append(userInfoList, SqlQuery(id, w))
		}(userId, &wg)
	}
	wg.Wait()
	fmt.Println(userInfoList)
}

func SqlQuery(userId int, wg *sync.WaitGroup) *UserInfo {
	defer wg.Done()
	u := &UserInfo{userId, "name", 1}
	//panic("123")
	return u
}
