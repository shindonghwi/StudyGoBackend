package main

// 읽기, 쓰기 func가 동기화가 되지 않음.

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

func main() {
	// RWMutex: 쓰기 Lock, 쓰기 시도중에는 다른곳에서 값 읽기 불가, 읽기 / 쓰기 Lock 모두 방지
	// RMutex: 읽기 Lock, 읽기 시도중에 값 변경 방지 / 쓰기 Lock 방지

	runtime.GOMAXPROCS(runtime.NumCPU())

	data := 0
	mutex := new(sync.RWMutex)

	go func() {
		for i := 1; i < 10; i++ {
			mutex.Lock()
			data += 1
			fmt.Println("Write: ", data)
			time.Sleep(200 * time.Millisecond)
			mutex.Unlock()
		}
	}()

	go func() {
		for i := 1; i < 10; i++ {
			mutex.RLock()
			fmt.Println("Read1: ", data)
			time.Sleep(1 * time.Second)
			mutex.RUnlock()
		}
	}()

	go func() {
		for i := 1; i < 10; i++ {
			mutex.RLock()
			fmt.Println("Read2: ", data)
			time.Sleep(1 * time.Second)
			mutex.RUnlock()
		}
	}()

	time.Sleep(10 * time.Second)
}
