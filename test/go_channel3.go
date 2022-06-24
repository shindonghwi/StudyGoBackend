package main

import (
	"fmt"
)

func main() {

	ch := make(chan bool)
	cnt := 5

	go func() {
		for i := 0; i < cnt; i++ {
			ch <- true
			fmt.Println("Go Func: ", i)
		}
		close(ch)
	}()

	for i := range ch {
		fmt.Println("Go Main: ", i)
	}

}
