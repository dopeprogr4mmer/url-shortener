package main

import (
	"encoding/json"
	"log"
	"net/http"
)

type ShortenRequest struct {
	URL string `json:"url"`
}

type ShortenResponse struct {
	ShortURL string `json:"short_url"`
}

func ShortenHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var req ShortenRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		log.Printf("[ERROR] JSON decode failed: %v", err)
		http.Error(w, "Bad request", http.StatusBadRequest)
		return
	}

	if req.URL == "" {
		log.Println("[WARN] Empty URL received")
		http.Error(w, "URL is required", http.StatusBadRequest)
		return
	}

	code := GenerateShortCode(req.URL)
	err := SaveURL(code, req.URL)
	if err != nil {
		log.Printf("[ERROR] SaveURL failed: %v", err)
		http.Error(w, "Server error", http.StatusInternalServerError)
		return
	}

	log.Printf("[INFO] Shortened URL: %s => %s", req.URL, code)

	resp := ShortenResponse{
		ShortURL: r.Host + "/" + code,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}

func RedirectHandler(w http.ResponseWriter, r *http.Request) {
	code := r.URL.Path[1:]
	if code == "" {
		http.NotFound(w, r)
		return
	}

	url, err := GetURL(code)
	if err != nil {
		log.Printf("[WARN] Unknown code: %s", code)
		http.NotFound(w, r)
		return
	}

	log.Printf("[INFO] Redirecting %s => %s", code, url)
	http.Redirect(w, r, url, http.StatusFound)
}
