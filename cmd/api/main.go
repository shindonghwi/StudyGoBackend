package main

import (
	"backend/models"
	"database/sql"
	"flag" // flag는 cmd에서 옵션을 지정할 수 있는 방법 중의 하나임.
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	_ "github.com/lib/pq"
)

const version string = "1.0.0"

type config struct {
	port int
	env  string
	db   struct {
		dsn string
	}
}

type AppStatus struct {
	Status      string `json:"status"`      // 상태
	Environment string `json:"environment"` // 환경 ( dev, prod )
	Version     string `json:"version" `    // 현재 버전
}

type application struct {
	config config
	logger *log.Logger
}

func main() {

	var cfg config

	flag.IntVar(&cfg.port, "port", 4000, "Server port to listen on")
	flag.StringVar(&cfg.env, "env", "development", "Application environment (development|production")
	flag.Parse()

	logger := log.New(os.Stdout, "", log.Ldate|log.Ltime)

	app := &application{
		config: cfg,
		logger: logger,
	}

	app.routes()

	srv := &http.Server{
		Addr:         fmt.Sprintf(":%d", cfg.port),
		Handler:      app.routes(),
		IdleTimeout:  time.Minute,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
	}

	logger.Println("Starting server on port", cfg.port)

	err := srv.ListenAndServe()

	if err != nil {
		log.Println(err)
	}

	db, err := sql.Open("mysql", "ehdgnl8940:ehdgnl8940!@tcp(52.12.181.219:3306)/Wolf")

	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	var movie models.Movie
	rows, err := db.Query("SELECT * FROM Movie")

	for rows.Next() {
		err := rows.Scan(
			&movie.ID,
			&movie.Title,
			&movie.Description,
			&movie.Year,
			&movie.ReleaseDate,
			&movie.Rating,
			&movie.Runtime,
			&movie.MPAARating,
			&movie.CreatedAt,
			&movie.UpdatedAt,
		)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(movie)
	}
}
