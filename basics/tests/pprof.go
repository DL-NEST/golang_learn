package main

import (
	"log"
	"net/http"
	_ "net/http/pprof"
)

func main() {
	mmser := make(map[int]int)

	for i := 0; i < 1000009000; i++ {
		mmser[i] = i
	}

	log.Println(http.ListenAndServe("localhost:6061", nil))
}
