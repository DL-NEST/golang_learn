package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan)
	go func() {
		for {
			time.Sleep(1 * time.Second)
			proc, _ := os.FindProcess(os.Getpid())
			err := proc.Signal(syscall.SIGINT)
			if err != nil {
				fmt.Printf("发送失败\n")
			}
		}
	}()
	s := <-signalChan
	fmt.Printf("程序结束%d", s)

	switch s {
	case syscall.SIGINT, syscall.SIGTERM:
		// 处理程序停止运行后的操作
		if o, ok := s.(syscall.Signal); ok {
			os.Exit(int(o))
		} else {
			os.Exit(0)
		}
	}
}
