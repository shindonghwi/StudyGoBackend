package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

func main() {

	runtime.GOMAXPROCS(runtime.NumCPU())

	var mutex = new(sync.Mutex)
	var condition = sync.NewCond(mutex)

	c := make(chan int, 5)

	for i := 0; i < 5; i++ {
		go func(n int) {
			mutex.Lock()
			c <- 777
			fmt.Println("Goroutine Wating, ", n)
			condition.Wait()
			fmt.Println("Wating End, ", n)
			mutex.Unlock()
		}(i)
	}

	for i := 0; i < 5; i++ {
		fmt.Println("receive: ", <-c)
	}

	for i := 0; i < 5; i++ {
		mutex.Lock()
		fmt.Println("Wake Goroutine: ", i)
		condition.Signal()
		mutex.Unlock()
	}

	mutex.Lock()
	condition.Broadcast()
	mutex.Unlock()

	time.Sleep(3 * time.Second)
}
