package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/sriganeshres/url_shortener/db"
	"github.com/sriganeshres/url_shortener/utils"
)

type ShortenRequest struct {
	URL         string `json:"url"`
	CustomAlias string `json:"customAlias"`
}

func ShortenURL(w http.ResponseWriter, r *http.Request) {
	var req ShortenRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil || req.URL == "" {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": "Invalid request"})
		return
	}

	shortCode := req.CustomAlias
	if shortCode == "" {
		shortCode = utils.GenerateShortCode()
	}

	err = db.SaveURL(req.URL, shortCode, req.CustomAlias != "")
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"error": "Could not save URL"})
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{
		"shortUrl": "http://localhost:8080/" + shortCode,
	})
}

