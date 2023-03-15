package main

import (
	"fmt"
	"github.com/panjf2000/ants/v2"
	"time"
)

type MyLogger struct {
}

func (a MyLogger) Printf(format string, args ...interface{}) {
	fmt.Printf("自定义logger")
	fmt.Printf(format, args)
}

func main() {
	defer ants.Release()

	//// 提交任务到ants包默认的协程池
	//err := ants.Submit(func() {
	//	for {
	//		time.Sleep(2 * time.Second)
	//		fmt.Println("well")
	//	}
	//})
	//if err != nil {
	//	return
	//}

	p, _ := ants.NewPool(5,
		// 预分配协程队列的内存
		ants.WithPreAlloc(true),
		// 在池子提交满了没有协程释放的时候提交阻塞，默认是阻塞的，使用非阻塞的时候当池子提交满了就会返回err
		ants.WithNonblocking(true),
	) //ants.WithLogger(MyLogger{}),

	for i := 0; i < 6; i++ {
		err := p.Submit(func() {
			for {
				time.Sleep(2 * time.Second)
				fmt.Println("well")
			}
		})
		if err != nil {
			fmt.Printf("err%d\n", i)
		}
	}
	fmt.Printf("running：%d\n", p.Running())
	for {

	}
}
