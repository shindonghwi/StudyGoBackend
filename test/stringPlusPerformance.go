package main

import (
	"fmt"
	"strings"
	"time"
)

func main() {

	myString1 := "The Go programming language is an open source project to make programmers more productive.\n" +
		"Go is expressive, concise, clean, and efficient. " +
		"Its concurrency mechanisms make it easy to write programs that get the most out of multicore and networked machines, " +
		"while its novel type system enables flexible and modular program construction. " +
		"Go compiles quickly to machine code yet has the convenience of garbage collection and the power of run-time reflection. " +
		"It's a fast, statically typed, compiled language that feels like a dynamically typed, interpreted language."

	myString2 := "Introduces the basics of creating and using multi-module workspaces in Go. " +
		"Multi-module workspaces are useful for making changes across multiple modules."

	elapsTime1, resultString1 := TestStringCombine1(myString1, myString2) // string + string
	elapsTime2, resultString2 := TestStringCombine2(myString1, myString2) // string.

	fmt.Printf("1번 실행시간: %s / 바이트 수 %d", elapsTime1, len(resultString1))
	fmt.Println()
	fmt.Printf("2번 실행시간: %s / 바이트 수 %d", elapsTime2, len(strings.Join(resultString2, "")))
}

func TestStringCombine1(str1 string, str2 string) (time.Duration, string) {
	startTime := time.Now()
	var tempStr string
	for i := 0; i < 10000; i++ {
		tempStr += str1 + str2
	}
	elapsedTime := time.Since(startTime)
	return elapsedTime, tempStr
}

func TestStringCombine2(str1 string, str2 string) (time.Duration, []string) {
	startTime := time.Now()
	var myStringSet []string
	for i := 0; i < 10000; i++ {
		myStringSet = append(myStringSet, str1)
		myStringSet = append(myStringSet, str2)
	}
	elapsedTime := time.Since(startTime)
	return elapsedTime, myStringSet
}
