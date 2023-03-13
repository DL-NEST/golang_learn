package main

import (
	"fmt"
	"golang_learn/utils"
)

type Cat struct {
	name string
}

func main() {
	fmt.Printf("slice")
	// 通过自建函数创建一个切片数组
	var sille = make([]string, 5)
	// 通过字面量创建一个切片数组
	_ = []int{1, 2, 3}
	// 通过下标创建一个切片数组
	_ = sille[:4]

	// 无法通过sille[5] = "ss" 来进行扩容,也会报错
	// 只能通过 append进行插入扩容
	// 在数据cap分配的空间不够时扩容会之间翻倍
	for i := 0; i < 10; i++ {
		sille = append(sille, "data")
		utils.OutputInfo(fmt.Sprintf("sille - for%d", i), sille, len(sille), cap(sille))
	}

	utils.OutputInfo("sille - 2", sille, len(sille), cap(sille))

	sille1 := make([]string, 2, 8)
	utils.OutputInfo("sille - 3", sille1, len(sille1), cap(sille1))
	sille1[1] = "dd"
	utils.OutputInfo("sille - 4", sille1, len(sille1), cap(sille1))

	silleSub := sille[2:5]
	utils.OutputInfo("sille - 5", silleSub, len(silleSub), cap(silleSub))

	// 基础库有copy可以对切片进行拷贝,当初始大小为0时copy函数会失效
	// 它只能用于切片，不能用于 map 等任何其他类型
	// 它返回结果为一个 int 型值，表示 copy 的长度
	// copy不会对切片进行动态扩容
	var protoSlice = []string{
		"date", "date",
	}
	var copyArray = make([]string, 0)
	var copyArray1 = make([]string, 10)
	a := copy(copyArray, protoSlice)
	fmt.Println(a) // 0
	utils.OutputInfo("copy - 1", copyArray, len(copyArray), cap(copyArray))
	b := copy(copyArray1, protoSlice)
	fmt.Println(b) // 2
	utils.OutputInfo("copy - 2", copyArray1, len(copyArray1), cap(copyArray1))

	// copy在元素为引用的时候只会拷贝引用地址
	matA := [][]int{
		{0, 1, 1, 0},
		{0, 1, 1, 1},
		{1, 1, 1, 0},
	}
	matB := make([][]int, len(matA))
	copy(matB, matA)
	fmt.Printf("%p, %p\n", matA, matA[0]) // 0xc0000c0000, 0xc0000c2000
	fmt.Printf("%p, %p\n", matB, matB[0]) // 0xc0000c0050, 0xc0000c2000

	matA1 := []*Cat{
		&Cat{
			name: "sss",
		},
	}
	matB1 := make([]*Cat, len(matA1))
	copy(matB1, matA1)
	fmt.Printf("%p, %p\n", matA1, &matA1[0].name) // 0xc0000c0000, 0xc0000c2000
	fmt.Printf("%p, %p\n", matB1, &matB1[0].name) // 0xc0000c0050, 0xc0000c2000
}
