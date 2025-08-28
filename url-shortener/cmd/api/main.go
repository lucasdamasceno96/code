package main

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/lucasdamasceno96/code/url-shortener/internal/config"
	"github.com/lucasdamasceno96/code/url-shortener/internal/handler"
	"github.com/lucasdamasceno96/code/url-shortener/internal/repository/sqlite"
	"github.com/lucasdamasceno96/code/url-shortener/internal/service"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	// Init config
	cfg := config.Config{
		ServerPort:   "8080",
		BaseURL:      "http://localhost:8080",
		FixedLetters: "k8s-",
		DBPath:       "./data/shortner.db",
	}

	// Connect to DB SQLite
	db, err := sql.Open("sqlite3", cfg.DBPath)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	//Init layers (Dependecy of injection)
	repo := sqlite.NewSQLiteRepo(db)
	svc := service.NewURLService(repo, cfg)
	h := handler.NewURLHandler(svc)

	// http routes
	r := mux.NewRouter()
	r.HandleFunc("/shorten", h.ShortenURL).Methods("POST")
	r.HandleFunc("/{code}", h.Redirect).Methods("GET")
	r.HandleFunc("/health", handler.Health).Methods("GET")

	log.Println("Server running on port", cfg.ServerPort)
	if err := http.ListenAndServe(":"+cfg.ServerPort, r); err != nil {
		log.Fatal(err)
	}
}
