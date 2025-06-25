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
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}

	shortCode := req.CustomAlias
	if shortCode == "" {
		shortCode = utils.GenerateShortCode()
	}

	err = db.SaveURL(req.URL, shortCode, req.CustomAlias != "")
	if err != nil {
		http.Error(w, "Could not save URL", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(map[string]string{
		"shortUrl": "http://localhost:8080/" + shortCode,
	})
}
