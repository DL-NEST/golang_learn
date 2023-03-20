package main

import (
	"fmt"
	"github.com/go-mmap/mmap"
	"io"
)

const filename = "./dex"

func main() {
	r, err := mmap.OpenFile(filename, mmap.Read|mmap.Write)
	if err != nil {
		fmt.Printf("Open: %+v\n", err)
	}
	defer r.Close()

	_, err = r.Stat()
	if err != nil {
		fmt.Printf("could not stat file: %+v\n", err)
	}

	got := make([]byte, r.Len())
	if _, err := r.ReadAt(got, 0); err != nil && err != io.EOF {
		fmt.Printf("ReadAt: %v\n", err)
	}
	fmt.Printf("%+v\n", got)
	fmt.Printf("%s\n", got)

	_, err = r.Write([]byte("c323232"))
	if err != nil {
		fmt.Printf("could not write-at: %+v\n", err)
	}
	//err = r.Sync()
	//if err != nil {
	//	fmt.Printf("could not sync mmap file: %+v\n", err)
	//}
	for {
		//err1 := r.Sync()
		//if err1 != nil {
		//	fmt.Printf("could not sync mmap file: %+v\n", err)
		//}
	}
}

//kafka-console-producer.sh --broker-list 127.0.0.1:9092 --topic first
//kafka-topics.sh --bootstrap-server 127.0.0.1:9092 --create --partitions 3 --replication-factor 1 --topic test
