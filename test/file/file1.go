package main

import (
	"fmt"
	"io/ioutil"
)

func errCheck(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {

	//s := "Hello Golang\nFile Write Test"

	//err := ioutil.WriteFile("text_test.txt", []byte(s), os.FileMode(0644))
	//errCheck(err)

	data, err := ioutil.ReadFile("text_test.txt")
	errCheck(err)

	fmt.Println(string(data))

}
