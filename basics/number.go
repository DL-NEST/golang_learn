package main

import (
	"fmt"
	"github.com/shopspring/decimal"
	"strconv"
)

func main() {
	fmt.Printf("number\n")

	// 有符号整型
	var _ int
	var _ int8  // 有符号 8 位整型 (-128 到 127)
	var _ int16 // -32768 到 32767
	var _ int32 // -2147483648 到 2147483647
	var _ int64 // -9223372036854775808 到 9223372036854775807

	// 无符号整型
	var _ uint
	var _ int8  // 0 到 255
	var _ int16 // 0 到 65535
	var _ int32 // 0 到 4294967295
	var _ int64 // 0 到 18446744073709551615

	var _ uintptr // 一种无符号整数类型，其大小足以容纳任何指针的位模式，主要是标志指针地址。
	var _ rune    // int32

	// 字符串转整型
	var numStr = "7749"
	num, err := strconv.Atoi(numStr)
	if err != nil {
		fmt.Printf("转换错误")
		return
	}
	fmt.Printf("%d\n", num)

	// 浮点型
	var _ float32 //
	var _ float64

	// github.com/shopspring/decimal 库解决浮点运算精度问题
	f1 := decimal.NewFromFloat(3.566)
	f2 := decimal.NewFromFloat(3.899)

	fmt.Printf("%f\n", f1.Add(f2).InexactFloat64())
}
