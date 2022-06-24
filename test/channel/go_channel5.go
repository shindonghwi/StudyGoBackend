package main

import (
	"fmt"
	"time"
)

func main() {

	intChannel := make(chan int)
	strChannel := make(chan string)

	go func() {
		for {
			intChannel <- 77
			time.Sleep(500 * time.Millisecond)
		}
	}()
	go func() {
		for {
			strChannel <- "Go lang"
			time.Sleep(250 * time.Millisecond)
		}
	}()

	go func() {
		for {
			select {
			case num := <-intChannel:
				fmt.Println("Int Data: ", num)
			case str := <-strChannel:
				fmt.Println("String Data: ", str)
			}
		}
	}()

	time.Sleep(7 * time.Second)
}
