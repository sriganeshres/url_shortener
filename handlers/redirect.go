package handlers

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/sriganeshres/url_shortener/db"
	"github.com/sriganeshres/url_shortener/kafka"
)

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
