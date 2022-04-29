package main

import (
	"bufio"
	//"bytes"
	"fmt"
	//"io"
	"net"
	"strconv"
	"time"
)

/**
func Listen(net, laddr string) (Listener, error)	获取监听器
func Accept()	监听器开始接收请求
func Dial(network, address string) (Conn, error)	向指定网络发送连接建立申请
func DialTimeout(network, address string, timeout time.Duration) (Conn, error)	申请连接并设定连接超时等待时间
func Read(b []byte) (n int, err error)	从socket接收缓冲区读数据
func Write(b []byte) (n int, err error)	向socket写入缓冲区写数据
func Close() (err error)	关闭连接
func LocalAddr() net.Addr | func RemoteAddr() net.Addr	获取当前连接的本地地址和远程地址
上面两方法的返回值可以调用两个方法：	.String() 网络地址 .Network() 协议类型
func SetDeadline(time.Time) error | SetReadDeadline| SetWriteDeadline	设置连接读、写超时时间
*/

func main() {
	// 将server端起来
	ch := make(chan bool)
	go serverStart(ch)

	serverState := <-ch
	fmt.Println("client find server state:", serverState, ",can goto start client......")

	// 客户端开启发送，注意：客户端可以不绑定(bind)本机地址，操作系统内核会为其分配一个
	//协议名 发送地址
	connClient, err := net.Dial("tcp", "127.0.0.1:8085")
	fmt.Println("client dial err:", err)

	//参数分别是： 协议名 发送地址  TCP连接超时时间
	//conn2, err := net.DialTimeout("tcp", "127.0.0.1:8085", 2*time.Second)
	//*************************************************************************************

	//第一种写入方式：通过write写入
	for i := 0; i < 10; i++ {
		connClient.Write([]byte("write:" + strconv.Itoa(i) + "\n"))
		fmt.Println("client write:" + strconv.Itoa(i) + " success")
		time.Sleep(time.Second)
	}

	//第二种写入方式：通过bufio写入 // bufio 包实现了带缓存的 I/O 操作
	/* writer := bufio.NewWriter(connClient)
	for i := 0; i < 10; i++ {
		writer.WriteString("write:" + strconv.Itoa(i))
		writer.WriteByte('\n')
		writer.Flush()
		fmt.Println("client write:" + strconv.Itoa(i) + " success")
		time.Sleep(time.Second)
	} */

	//select{}
}

func serverStart(ch chan bool) {
	// 服务器端开启监听
	listener, err := net.Listen("tcp", "127.0.0.1:8085") //协议名 监听地址
	fmt.Printf("server net listen:%v err:%v\n", listener.Addr().String(), err)
	ch <- true

	conn, err := listener.Accept()
	//*************************************************************************************
	//获取协议类型
	fmt.Println("server revd conn network type:", conn.RemoteAddr().Network())
	//获取网络地址
	fmt.Println("server revd conn addr:", conn.RemoteAddr().String())
	//*************************************************************************************

	/* for {
		//1.第一种读取方式：直接通过Buffer读取
		var dataBuffer bytes.Buffer
		b := make([]byte, 10)
		n, err := conn.Read(b)
		if err != nil {
			fmt.Println("server read data err:", err)
			if err == io.EOF { //发现TCP连接在另一端已经关闭
				conn.Close()
			} else {
				fmt.Printf("error:%s\n", err)
			}
			break
		}
		dataBuffer.Write(b[:n])
		fmt.Println("server revd:",dataBuffer.String())

		// SetDeadline设置的是一个绝对时间（即具体时间），并且会对以后的每次读写都生效
		// 所以循环读写的操作，需要不断更新这个时间
		// b := make([]byte, 10)
		// for {
		// 	conn.SetDeadline(time.Now().Add(2 * time.Second))
		// 	n, err := conn.Read(b)
		// }
		// conn.SetDeadline(time.Time{}) //取消超时时间设置
	} */

	for {
		//第二种读取方式：2.通过bufio读取 // bufio 包实现了带缓存的 I/O 操作
		reader := bufio.NewReader(conn)
		line, err := reader.ReadBytes('\n')
		if err != nil {
			fmt.Println("server bufio read err:", err)
			continue
		}
		fmt.Println("server revd:", string(line))
	}

}
