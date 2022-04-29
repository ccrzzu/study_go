package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
)

/**
func Notify(c chan<- os.Signal, sig …os.Signal)	当操作系统向当前进程发送指定信号时发出通知
func Stop(c chan<- os.Signal)	删除定义的自处理通知通道，恢复系统默认操作
func os.FindProcess()(pid int) (*Process, error)	根据Pid返回一个进程对象，可以向对象发送Signal信息
syscall.SIGINT…	syscall中定义了所有信号常量
*/

func main() {
	/**
		1. Notify不会因为sigRecv1满而被阻塞
		2. 接收到信号而不做合适的处理，相当于程序忽略掉了系统发来的信号
	    3. SIGKILL和SIGSTOP永远不会被自处理或者忽略（因为他们是提供给超级用户终止或停止进程的可靠方法）
		即使Notify(,syscall.SIGKILL, syscall.SIGSTOP)也不会改变系统的处理动作
	*/
	//同一个进程可以建立多个自定义信号处理通道
	sigRecv1 := make(chan os.Signal, 1)
	sigs1 := []os.Signal{syscall.SIGINT, syscall.SIGQUIT}
	fmt.Printf("Set notification for %s... [sigRecv1]\n", sigs1)

	sigRecv2 := make(chan os.Signal, 1)
	sigs2 := []os.Signal{syscall.SIGINT}
	fmt.Printf("Set notification for %s... [sigRecv2]\n", sigs2)

	// arg1：监测到信号时由此通知  arg2：监测的信号类型（不填则监控所有类型）
	signal.Notify(sigRecv1, sigs1...)
	signal.Notify(sigRecv2, sigs2...)

	// 删除sigRecv1，之后恢复以系统默认方式处理信号
	signal.Stop(sigRecv1)
	// 此时只剩下sigRecv2还在工作

	for {
		switch <-sigRecv2 {
		case syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT:
			fmt.Println("revd signal self handle......")
			return
		}
	}
}

//rtc center的例子
func Loop(Quit, Hup func()) {
	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGHUP, syscall.SIGQUIT,
		syscall.SIGTERM, syscall.SIGINT, syscall.SIGTRAP)
	for {
		switch <-c {
		case syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT:
			if Quit == nil {
				return
			}
			Quit()
			return
		case syscall.SIGHUP:
			if Hup == nil {
				return
			}
			Hup()
			return
		}
	}
}
