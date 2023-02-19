package main

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/go-redis/redis/v8"
	"time"
)

type Person struct {
	Name    string `json:"name,omitempty"`
	Moniker string `json:"moniker"`
	// 小写字段和json标记为-的字段不导出
	Gender bool  `json:"-"`
	Other  Other `json:"other"`
}

type Other struct {
	Tel int `json:"tel,omitempty"`
	// 默认编码为 base64 编码的字符串
	HeadImg []byte `json:"head_img,omitempty"`
}

func main() {

	RdbAuth := redis.NewClient(&redis.Options{
		Addr:     "127.0.0.1:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	var person = Person{
		Name:    "ues",
		Moniker: "none",
		Gender:  false,
		Other: Other{
			Tel: 23333333,
			HeadImg: []byte{
				'A', 'B', 'C',
			},
		},
	}

	var personList = []Person{
		person,
		person,
	}
	// 对象转换为JSON字符串
	// 数组和切片值编码为 JSON 数组，但 []byte 编码为 base64 编码的字符串，而 nil 切片编码为 null JSON 值。
	// 通道值、复数值和函数值不能用 JSON 编码。尝试对此类值进行编码会导致 Marshal 返回 UnsupportedTypeError。
	// 布尔值编码为 JSON 布尔值。浮点、整数和数字值编码为 JSON 数字。字符串值编码为强制为有效 UTF-8 的 JSON 字符串，将无效字节替换为 Unicode 替换符文。
	// 数组和切片值编码为 JSON 数组，但 []byte 编码为 base64 编码的字符串，而 nil 切片编码为 null JSON 值。
	res, _ := json.Marshal(personList)

	fmt.Printf("Marshal: %s\n", res)

	// 存redis
	RdbAuth.Set(context.Background(), "test", res, 100*time.Second)
	s, _ := RdbAuth.Get(context.Background(), "test").Bytes()

	var UnPersonList []Person

	err := json.Unmarshal(s, &UnPersonList)
	if err != nil {
		fmt.Printf("erro\n")
		return
	}
	fmt.Printf("%v\n", UnPersonList)
	fmt.Printf("%v\n", string(UnPersonList[0].Other.HeadImg))

}
