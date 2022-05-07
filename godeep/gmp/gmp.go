package main

import (
	"fmt"
	"os"
	"runtime/trace"
	"time"
)

/**
可视化查看GMP信息的方式
第一种：
运行程序 go run gmp.go
会得到一个trace.out文件，然后我们可以用一个工具打开，来分析这个文件。
再执行 go tool trace trace.out
output:
Opening browser. Trace viewer is listening on http://127.0.0.1:33479
我们可以通过浏览器打开http://127.0.0.1:33479网址，点击view trace 能够看见可视化的调度流程。
*/

func main1() {

	//创建trace文件
	f, err := os.Create("trace.out")
	if err != nil {
		panic(err)
	}

	defer f.Close()

	//启动trace goroutine
	err = trace.Start(f)
	if err != nil {
		panic(err)
	}
	defer trace.Stop()

	//main
	fmt.Println("Hello World")
}

/**
第二种：
go build gmp.go
GODEBUG=schedtrace=1000 ./gmp 
*/
func main() {
	for i := 0; i < 5; i++ {
		time.Sleep(time.Second)
		fmt.Println("Hello World")
	}
}
