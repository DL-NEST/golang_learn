package main

import (
	"context"
	"fmt"
	"time"
)

func main() {

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)

	defer cancel()

	go func() {
		fmt.Println("等待")
		select {
		case <-ctx.Done():
			fmt.Println("结束")
			return
		}
	}()
	for {

	}
}
