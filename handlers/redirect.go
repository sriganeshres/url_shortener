package handlers

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/sriganeshres/url_shortener/db"
	"github.com/sriganeshres/url_shortener/kafka"
)

/*
 * @brief RedirectURL handles GET requests to redirect from short URLs to original URLs.
 * 
 * This function retrieves the original URL corresponding to a given short code from the request.
 * If the short code is found, the request is redirected to the original URL. Otherwise, it responds
 * with an error message indicating the URL was not found. Additionally, each visit is asynchronously
 * logged to Kafka.
 * 
 * @param w The ResponseWriter to write HTTP responses.
 * @param r The incoming HTTP request containing the short code.
 */
func RedirectURL(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	shortCode := vars["shortCode"]

	originalURL, err := db.GetOriginalURL(shortCode)
	if err != nil {
		http.Error(w, "URL not found", http.StatusNotFound)
		return
	}

	go kafka.LogVisit(shortCode, r.RemoteAddr)

	http.Redirect(w, r, originalURL, http.StatusFound)
}
