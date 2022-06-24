package main

import (
	"fmt"
	"time"
)

func main() {
	/**
	채널은 기본적으로 동기적으로 동작한다.
	(채널 <- 데이터) 를 통해서 송신하고
	(변수 <- 채널 or <- 채널) 을 통해 수신한다.

	송신을 하는 쪽과 수신을 하는쪽이 suspension point 가 되며
	suspension point 아래에 있는 코드는 반대쪽 suspension point가 실행되기 전까지 잠시 잠들어있는다.

	*/

	ch := make(chan bool)
	cnt := 12

	go func() {
		for i := 0; i < cnt; i++ {
			ch <- true
			fmt.Println("Go Func: ", i)
			time.Sleep(1 * time.Second)
		}
	}()

	for i := 0; i < cnt; i++ {
		<-ch
		fmt.Println("Main: ", i)
	}

}
