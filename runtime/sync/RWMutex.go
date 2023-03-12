package main

import (
	"fmt"
	"os"
	"sync"
)

func main() {
	// RWMutex是基于Mutex的
	// 读写互斥锁在互斥锁之上提供了额外的更细粒度的控制，能够在读操作远远多于写操作时提升性能。
	lock := sync.RWMutex{}
	open, err := os.OpenFile("./file", os.O_RDWR|os.O_CREATE|os.O_APPEND, 777)
	if err != nil {
		fmt.Println("文件打开失败")
		return
	}

	go func() {
		// 写锁
		lock.Lock()
		_, errWrite := open.Write([]byte("file\n"))
		if errWrite != nil {
			fmt.Println("文件写入失败")
			return
		}
		// 释放写锁
		lock.Unlock()
	}()

	for i := 0; i < 10; i++ {
		ii := i
		go func() {
			// 写锁
			lock.Lock()
			_, errWrite := open.Write([]byte(fmt.Sprintf("file%d\n", ii)))
			if errWrite != nil {
				fmt.Println("文件写入失败")
				return
			}
			// 释放写锁
			lock.Unlock()
		}()
	}

	for j := 0; j < 10; j++ {
		go func() {
			// 写锁
			lock.RLock()
			var data []byte
			_, errWrite := open.Read(data)
			if errWrite != nil {
				fmt.Println("文件读取失败")
				return
			}
			fmt.Printf("%v", data)
			// 释放写锁
			lock.RUnlock()
		}()
	}
	defer open.Close()

	for {

	}

}
