package main

// #include <stdio.h>
// void callC() {
//   printf("Hello World from C!\n");
// }
import "C"
import "fmt"

func main() {

	fmt.Println("让我们学习 Go 语句调用 C 程序")
	C.callC()
	fmt.Println("调用 C 程序结束")
}
