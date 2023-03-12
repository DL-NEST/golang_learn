package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	// 互斥锁
	// 比较重的锁
	// 加锁过程比较复杂
	lock := sync.Mutex{}

	go func() {
		lock.Lock()
		fmt.Println("加锁了的")

		time.Sleep(3 * time.Second)
		lock.Unlock()
	}()

	for i := 0; i < 10; i++ {
		ii := i
		go func() {
			lock.Lock()
			fmt.Printf("获得锁后%d", ii)
			time.Sleep(1 * time.Second)
			lock.Unlock()
		}()
	}

	for {

	}
}
