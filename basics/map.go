package main

import (
	"fmt"
	"golang_learn/utils"
)

func main() {
	fmt.Printf("map")
	var mp = make(map[string]string)

	utils.OutputInfo("map", mp, len(mp), 3)
}
