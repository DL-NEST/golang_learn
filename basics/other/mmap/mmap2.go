package main

import (
	"fmt"
	mp "github.com/edsrzf/mmap-go"
	"os"
)

func main() {
	//fmt.Println()
	f, err := os.OpenFile("./dex.log", os.O_RDWR|os.O_CREATE|os.O_SYNC, 0755)
	if err != nil {
		println("打开失败")
	}
	m, err := mp.MapRegion(f, 4096*256*5, mp.RDWR, 0, 4096*13)
	//m, err := mp.Map(f, mp.COPY, 0)
	if err != nil {
		println("映射失败")
		println(err.Error())
		return
	}
	fmt.Println(string(m))
	fmt.Printf("%d\n", len(m))
	m[1] = 1
}
