package main

import "fmt"

func receiveOnly(cnt int) <-chan int {

	sum := 0
	tot := make(chan int)

	go func() {
		for i := 0; i < cnt; i++ {
			sum += i
		}
		tot <- sum
		tot <- 777
		tot <- 888
		close(tot)
	}()

	return tot
}

func total(c <-chan int) <-chan int {
	sum := 0
	tot := make(chan int)

	go func() {
		for i := range c {
			sum += i
		}
		tot <- sum
	}()

	return tot
}

func main() {

	c := receiveOnly(100)
	output := total(c)

	fmt.Println("Main: ", <-output)
}
