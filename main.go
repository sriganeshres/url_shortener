package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/sriganeshres/url_shortener/db"
	"github.com/sriganeshres/url_shortener/handlers"
	"github.com/sriganeshres/url_shortener/kafka"
)

func main() {
	db.InitPostgres()
	db.InitRedis()
	kafka.InitKafka()

	r := mux.NewRouter()
	r.HandleFunc("/shorten", handlers.ShortenURL).Methods("POST")
	r.HandleFunc("/{shortCode}", handlers.RedirectURL).Methods("GET")

	log.Println("Server running at :8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
