package main

import (
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Println("[WARN] No .env file found, using default config.")
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
		log.Println("[INFO] Defaulting to port 8080")
	}

	dbPath := os.Getenv("DB_PATH")
	if dbPath == "" {
		dbPath = "shortener.db"
		log.Println("[INFO] Using default DB path shortener.db")
	}

	err = InitDB(dbPath)
	if err != nil {
		log.Fatalf("[FATAL] DB init failed: %v", err)
	}
	log.Println("[INFO] Database initialized")

	http.HandleFunc("/shorten", ShortenHandler)
	http.HandleFunc("/", RedirectHandler)

	log.Printf("[INFO] Server starting at :%s 🚀", port)
	err = http.ListenAndServe(":"+port, nil)
	if err != nil {
		log.Fatalf("[FATAL] Server failed to start: %v", err)
	}
}
