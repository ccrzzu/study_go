package main

import (
	"context"
	"fmt"
	"time"

	"golang.org/x/time/rate"
)

func main() {
	r := rate.NewLimiter(1, 5) // 每秒放置1个令牌，最多存储5个令牌
	fmt.Println(r.Burst())     // 返回令牌桶的最大容量
	fmt.Println(r.Limit())     // 返回令牌桶的生成速率，也就是说每秒产生多少个令牌
	fmt.Println(r.Tokens())

	// Allow(l Limit) bool:
	//该方法用于检查是否可以处理一个事件（或一个令牌），返回一个布尔值。如果令牌桶中有足够的令牌以满足要求，则返回true；否则返回false。
	/* r := rate.NewLimiter(1, 5) // 每秒放置1个令牌，最多存储5个令牌
	for i := 0; i < 10; i++ {
		if r.Allow() {
			fmt.Println("Handle event", i)
		} else {
			fmt.Println("Rate limited", i)
		}
	} */

	//AllowN(l Limit, n int) bool:
	//该方法与Allow方法类似，但可以同时处理n个事件（或令牌），返回一个布尔值。如果令牌桶中有足够的令牌以满足要求，则返回true；否则返回false。
	/* for i := 0; i < 10; i++ {
		fmt.Println(r.Tokens())
		if r.AllowN(time.Now(), 2) { // 每次处理2个事件
			fmt.Println("Handle events", i, i+1)
		} else {
			fmt.Println("Rate limited", i, i+1)
		}
		time.Sleep(time.Second)
	} */

	// 	Reserve(l Limit) *Reservation:
	// 该方法返回一个Reservation对象。Reservation对象表示一个持续时间段，表示一个事件（或令牌）的预订。使用该对象可以等待预订的时间段，然后再处理事件（或令牌）。
	/* for i := 0; i < 10; i++ {
		res := r.Reserve() // 预订一个时间段
		if res.OK() {
			fmt.Println("Handle event", i)
			time.Sleep(res.Delay()) // 等待预订的时间段
		} else {
			fmt.Println("Rate limited", i)
		}
	} */

	// ReserveN(l Limit, n int) *Reservation:
	// 该方法与Reserve方法类似，但可以同时预订n个事件（或令牌）。返回一个Reservation对象。(具体不在重复，代码可参考Reserve)

	// SetBurst(b int):
	// 该方法用于动态设置令牌桶的容量，即最大的令牌数量。

	// SetBurstAt(t time.Time, b int):
	// 该方法用于在给定的时间设置令牌桶的容量。

	// SetLimit(l Limit):
	// 该方法用于动态设置令牌桶的限制速率。

	// SetLimitAt(t time.Time, l Limit):
	// 该方法用于在给定的时间设置令牌桶的限制速率。

	// Tokens() float64:
	// 该方法返回令牌桶中当前的令牌数量。

	// TokensAt(t time.Time) float64:
	// 该方法返回在给定的时间令牌桶中的令牌数量。

	//Wait(t time.Time) bool:
	//该方法等待直到指定的时间t再处理事件（或令牌）。返回一个布尔值，如果在指定的时间段内有足够的令牌，则返回true；否则返回false。

	for i := 0; i < 10; i++ {
		ctx, _ := context.WithDeadlineCause(context.Background(), time.Now().Add(100*time.Millisecond), fmt.Errorf("111 context deadline"))
		if err := r.Wait(ctx); err == nil { // 等待1秒钟
			fmt.Println("Handle event", i)
		} else {
			fmt.Println("Rate limited", i, ctx.Err())
		}
	}
}
