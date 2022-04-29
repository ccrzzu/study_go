package main

import (
	"bytes"
	"fmt"
	"os/exec"
)

/* type S struct {
}

func f(x interface{}) {
}

func g(x *interface{}) {
}

func main() {
	s := S{}
	p := &s
	f(s) //A
	g(s) //B
	f(p) //C
	g(p) //D
} */

func main() {
	cmd1 := exec.Command("ps", "aux")
	cmd2 := exec.Command("grep", "redis")

	//cmd对象的底层，Stdout与Stdin属性也是通过指向一个字节流实现读写的，这里用新建的字节流代替
	var outputBuf1 bytes.Buffer
	cmd1.Stdout = &outputBuf1
	if err := cmd1.Start(); err != nil {
		fmt.Printf("Error: The first command can not be startup %s\n", err)
		return
	}
	//wait会阻塞cmd直到其运行完毕
	if err := cmd1.Wait(); err != nil {
		fmt.Printf("Error: Couldn't wait for the first command: %s\n", err)
		return
	}

	//cmd1的输出与cmd2的输入指向同一个字节流地址，管道
	cmd2.Stdin = &outputBuf1
	var outputBuf2 bytes.Buffer
	cmd2.Stdout = &outputBuf2
	if err := cmd2.Start(); err != nil {
		fmt.Printf("Error: The second command can not be startup: %s\n", err)
		return
	}
	if err := cmd2.Wait(); err != nil {
		fmt.Printf("Error: Couldn't wait for the second command: %s\n", err)
		return
	}

	//打印cmd2的输出
	fmt.Println(cmd2.Stdout)
}
