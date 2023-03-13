package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	// 确保只会运行一次
	var once sync.Once

	go func() {
		// 传入一个函数，它确保只会运行一次
		once.Do(func() {
			fmt.Println("run once")
		})
	}()
	go func() {
		once.Do(func() {
			fmt.Println("run once")
		})
	}()

	time.Sleep(10 * time.Second)
}
