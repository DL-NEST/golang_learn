package main

import (
	"fmt"
	"regexp"
)

// 监听信号量
func main() {
	//compile, err := regexp.Compile("/init/*")
	//if err != nil {
	//	return
	//}
	matchString, err := regexp.MatchString("/init/*", "/init/ss")
	if err != nil {
		return
	}
	fmt.Printf("%v", matchString)
}
