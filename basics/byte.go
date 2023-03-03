package main

import (
	"fmt"
)

func main() {

	c := 'a'

	var aa byte

	aa = 0x04

	fmt.Printf("%b\n", c)
	fmt.Printf("%b\n", aa)

	if bin(aa, 3) {
		fmt.Printf("1")
	} else {
		fmt.Printf("0")
	}

}

func bin(bin byte, len uint) bool {
	if (bin & (1 << (len - 1))) > 0 {
		return true
	}
	return false
}
