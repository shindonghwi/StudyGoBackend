package main

import (
	"fmt"
	"runtime"
)

func main() {

	runtime.GOMAXPROCS(1)

	/**

	make(chan type, 버퍼 용량) 과 같이 버퍼 용량을 지정 할 수 있다. ( 비동기 )

	버퍼:
	송신 하는쪽이 가득차면 대기하고, 비어있으면 다시 작동한다.
	수신 하는쪽이 비어있으면 대기하고, 가득차면 다시 작동한다.
	*/

	ch := make(chan bool, 3)
	cnt := 12

	go func() {
		for i := 0; i < cnt; i++ {
			ch <- true
			fmt.Println("Go Func: ", i)
		}
	}()

	for i := 0; i < cnt; i++ {
		<-ch
		fmt.Println("Main: ", i)
	}

}
