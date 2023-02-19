package main

import (
	"fmt"
	"golang_learn/utils"
)

func main() {
	fmt.Printf("array")
	// 收集错误信息
	defer func() {
		if err := recover(); err != nil { //注意必须要判断
			fmt.Printf("发送错误： %s", err)
		}
	}()
	// 声明一个数组变量,需要声明数组存储的数据类型
	// 数组其实也是一种对象通过types.NewArray()初始化
	var array [5]string
	// ... 是golang的一种语法糖,使用时，golang会通过实际数据去判断数组的长度来给数组分配空间
	// 当数组的长度为4或及以下的时候会直接将数组中的元素放置在栈上；
	// 当元素数量大于4时，会将数组中的元素放置到静态区并在运行时取出
	var array1 = [...]string{
		"data", "data",
	}
	utils.OutputInfo("array init", array, len(array), cap(array))
	utils.OutputInfo("array1 init", array1, len(array1), cap(array1))

	// array[6] = "data" 会发出错误，数组类型不能动态扩容，当超出类型的大小的时候会发生数组越界程序会发出panic
}
