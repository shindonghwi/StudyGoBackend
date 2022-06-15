package main

import (
	"encoding/json"
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

type AppStatus struct {
	Status      string // 상태
	Environment string // 환경 ( dev, prod )
	Version     string // 현재 버전
}

func main() {

	var cfg config

	flag.IntVar(&cfg.port, "port", 4000, "Server port to listen on")
	flag.StringVar(&cfg.env, "env", "development", "Application environment (development|production")
	flag.Parse()

	fmt.Println("Running: ", cfg.port, cfg.env)

	// HandleFunc ( 수신할 경로, 핸들러 함수, http 포인터 )
	http.HandleFunc("/status", func(w http.ResponseWriter, r *http.Request) {
		// fmt.Fprint(w, "status")
		currentStatus := AppStatus{
			Status:      "Available",
			Environment: cfg.env,
			Version:     version,
		}

		js, err := json.MarshalIndent(currentStatus, "", "'\t")

		if err != nil {
			log.Println(err)
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(js)
	})

	if err := http.ListenAndServe(fmt.Sprintf(":%d", cfg.port), nil); err != nil {
		log.Println(err)
	}
}
