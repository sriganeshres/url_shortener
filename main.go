package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/sriganeshres/url_shortener/db"
	"github.com/sriganeshres/url_shortener/handlers"
	"github.com/sriganeshres/url_shortener/kafka"
)

/*
 * @brief main is the entry point for the URL shortener service.
 *
 * This function initializes the database, Redis, and Kafka connections and
 * sets up the HTTP routes for the service. The service listens on port 8080.
 */
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
