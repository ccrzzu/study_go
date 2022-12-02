package array

import (
	"math/rand"
	"time"
)

/**
念力云游戏：
基础架构
资源调度-云上
操作系统-虚拟
RTC-低延迟
*/
type User struct {
	ID  int //userid
	Val int //权重值
}

//根据user的权重值返回一个user
//例如，三个user权重分别是3，2，5
//那么他们的区间分别是[0,3),[3,5),[5,9)
//通过0到9的一个随机数
//根据这个随机数循环减掉这些user的权重值，得出是负值时就证明它落在相应的区间了
func RangeJudgeGetUser(userList []User) User {
	user := User{}
	sum := 0
	for _, item := range userList {
		sum += item.Val
	}

	// 初始化随机数的资源库, 如果不执行这行, 不管运行多少次都返回同样的值
	rand.Seed(time.Now().UnixNano())
	n := rand.Intn(sum)

	for _, item := range userList {
		if n-item.Val >= 0 {
			n = n - item.Val
		} else {
			user = item
		}
	}
	return user
}
