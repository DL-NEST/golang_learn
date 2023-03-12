package main

import (
	"fmt"
	"sync"
)

func main() {
	// 值不允许被拷贝
	// 等待一组goroutine返回
	var wg sync.WaitGroup
	var goList = 10

	wg.Add(goList)

	for i := 0; i < goList; i++ {
		ii := i
		go func() {
			// 结束后
			fmt.Printf("%d\n", ii)
			wg.Done()
		}()
	}
	// 阻塞等待WaitGroup全部done出去
	wg.Wait()
	fmt.Println("运行结束")
}
