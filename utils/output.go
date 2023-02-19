package utils

import "fmt"

func OutputInfo(name string, data any, len int, cap int) {
	fmt.Printf("title: %v\t", name)
	fmt.Printf("data: %v\t", data)
	// 数据长度
	fmt.Printf("len: %v\t", len)
	// 数据空间
	fmt.Printf("cap: %v\t\n", cap)
}
