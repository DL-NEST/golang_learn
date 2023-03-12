package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	// 带条件的锁
	// 等待父准备完毕后启动所有goroutine
	c := sync.NewCond(&sync.Mutex{})

	for i := 0; i < 10; i++ {
		go func() {
			c.L.Lock()
			c.Wait()
			fmt.Println("广播")
			c.L.Unlock()
		}()
	}

	go func() {
		time.Sleep(3 * time.Second)
		fmt.Println("发送")
		c.L.Lock()
		c.Broadcast()
		c.L.Unlock()
	}()
	for {

	}
}
