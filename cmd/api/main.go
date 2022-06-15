package main

import (
	"flag" // flag는 cmd에서 옵션을 지정할 수 있는 방법 중의 하나임.
	"fmt"
	"log"
	"net/http"
)

const version string = "1.0.0"

type config struct {
	port int
	env  string
}

func main() {

	var cfg config

	flag.IntVar(&cfg.port, "port", 4000, "Server port to listen on")
	flag.StringVar(&cfg.env, "env", "development", "Application environment (development|production")
	flag.Parse()

	fmt.Println("Running: ", cfg.port, cfg.env)

	http.HandleFunc("/status", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "status")
	})

	err := http.ListenAndServe(fmt.Sprintf(":%d", cfg.port), nil)

	if err != nil {
		log.Println(err)
	}
}
