package main

import (
	"fmt"
	"net"
)

func main() {
	// 创建一个tcp连接
	conn, err := net.Dial("tcp", "127.0.0.1:7789")
	if err != nil {
		fmt.Println("连接失败")
		return
	}

	_, err = conn.Write([]byte("hallo"))
	if err != nil {
		fmt.Println("发送失败")
		return
	}
	defer conn.Close()
}
