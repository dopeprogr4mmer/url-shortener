package main

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

var db *sql.DB

func InitDB(path string) error {
	var err error
	db, err = sql.Open("sqlite3", path)
	if err != nil {
		return err
	}

	query := `
	CREATE TABLE IF NOT EXISTS urls (
		code TEXT PRIMARY KEY,
		url TEXT NOT NULL
	);`
	_, err = db.Exec(query)
	if err != nil {
		return err
	}

	log.Println("[INFO] URL table ready")
	return nil
}

func SaveURL(code, url string) error {
	log.Printf("[DEBUG] Saving URL: %s => %s", code, url)
	_, err := db.Exec("INSERT OR REPLACE INTO urls(code, url) VALUES(?, ?)", code, url)
	return err
}

func GetURL(code string) (string, error) {
	var url string
	err := db.QueryRow("SELECT url FROM urls WHERE code = ?", code).Scan(&url)
	if err != nil {
		return "", err
	}
	return url, nil
}
