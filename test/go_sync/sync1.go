package main

// 공유데이터를 동기화 시키지 않았기 때문에 출력 결과값이 항상 달리진다.

import (
	"fmt"
	"runtime"
	"sync"
)

type count struct {
	cnt   int
	mutex sync.Mutex
}

func (c *count) increment() {
	c.mutex.Lock()
	c.cnt += 1
	c.mutex.Unlock()
}

func (c *count) result() {
	fmt.Println("done: ", c.cnt)
}

func main() {

	runtime.GOMAXPROCS(runtime.NumCPU())

	c := count{cnt: 0}
	done := make(chan bool)

	for i := 0; i < 10000; i++ {
		go func() {
			c.increment()
			done <- true
			runtime.Gosched() // CPU 양보한다.
		}()
	}

	for i := 0; i < 10000; i++ {
		<-done
	}

	c.result()
}
