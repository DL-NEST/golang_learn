package main

import (
	"context"
	"fmt"
	"time"
)

func main() {

	ctx, cancel := context.WithCancel(context.Background())
	ch := make(chan int)

	go func() {
		fmt.Println("等待")
		//time.Sleep(2 * time.Second)
		select {
		case <-ctx.Done():
			fmt.Println("结束")
		case val := <-ch:
			fmt.Printf("%d", val)
		}
	}()
	ch <- 32
	time.Sleep(2 * time.Second)
	cancel()
	//val <- 32
	time.Sleep(5 * time.Second)
}
