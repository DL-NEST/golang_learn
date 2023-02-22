package main

import (
	"fmt"
	"sync"
)

/* 懒汉模式 */
// 使用的时候才创建实列

// 不加锁全局调用会有线程问题
// 单例加锁

type logger struct {
	item int
}

var log *logger

// 原子性
var logLock sync.Once

func (l *logger) add() {
	l.item += 1
}

func GetLog() *logger {
	if log == nil {
		// 函数具有原子性，只会触发一次
		logLock.Do(func() {
			log = &logger{item: 0}
		})
	}
	return log
}

func main() {
	var l *logger

	l = GetLog()
	l.add()

	fmt.Printf("%d", l.item)
}
