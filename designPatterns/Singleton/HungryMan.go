package main

/* 饿汉模式 */
// 程序运行的时候就创建实列，不管用不用，可能会浪费一些空间，但是没有线程安全问题

var logg = &logger{item: 0}

func main() {
	logg.add()
}
