package main

import (
	"fmt"
	"time"
)

func main() {
	// go关键字可以创建goroutines，golang的运行最小单元，也就是协程
	go func() {
		// 协程在父goroutines终止时也会停止
		for {
			time.Sleep(1 * time.Second)
			fmt.Println("child")
		}
	}()

	go task()

	time.Sleep(20 * time.Second)
}

func task() {

}
