package main

import (
	"fmt"
	"time"
)

func main() {
	// chan构造
	//ch : make(chan struct{}) // 空类型做信号发送接收
	//ch := make(chan int)    // 无缓冲
	ch := make(chan int, 10) // 有缓冲

	// 对未初始化的channel读写操作都会死锁
	//var ch1 chan int
	// 死锁
	//ch1 <- 3

	// 关闭channel
	// 关闭channel的时候会唤醒所有被channel读阻塞的goroutines
	//close(ch)

	// 写channel
	ch <- 5
	go func() {
		// 读channel,第二个参数会读取channel的isClose判断channel是否被关闭
		if c, ok := <-ch; ok {
			fmt.Println(c)
			fmt.Println("channel未关闭")
		} else {
			fmt.Println("channel已关闭")
		}
	}()

	time.Sleep(10 * time.Second)
}
