package main

// 읽기, 쓰기 func가 동기화가 되지 않음.

import (
	"fmt"
	"runtime"
	"time"
)

func main() {

	runtime.GOMAXPROCS(runtime.NumCPU())

	data := 0

	go func() {
		for i := 1; i < 10; i++ {
			data += 1
			fmt.Println("Write: ", data)
			time.Sleep(200 * time.Millisecond)
		}
	}()

	go func() {
		for i := 1; i < 10; i++ {
			fmt.Println("Read1: ", data)
			time.Sleep(1 * time.Second)
		}
	}()

	go func() {
		for i := 1; i < 10; i++ {
			fmt.Println("Read2: ", data)
			time.Sleep(1 * time.Second)
		}
	}()

	time.Sleep(5 * time.Second)
}
