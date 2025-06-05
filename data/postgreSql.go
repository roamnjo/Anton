package data

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/lib/pq"
)

type Storage struct {
	db *sql.DB
}

func ConnectDB() (*Storage, error) {
	dbUrl := os.Getenv("DB_URL")

	connStr := dbUrl

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}

	stmt, err := db.Prepare(`
	CREATE TABLE IF NOT EXISTS url (
	id SERIAL PRIMARY KEY,
	url VARCHAR(200) NOT NULL,
	alias VARCHAR(64) NOT NULL)
	`)

	defer stmt.Close()

	if err != nil {
		log.Println(err)
	}

	_, err = stmt.Exec()
	if err != nil {
		log.Println(err)
	}

	return &Storage{db: db}, nil

}

func (s *Storage) SaveURL(url, alias string) (int64, error) {
	var id int64

	err := s.db.QueryRow("INSERT INTO url (url, alias) VALUES($1, $2) RETURNING id", url, alias).Scan(&id)
	if err != nil {
		return 0, fmt.Errorf("SaveURL: save data failed %w", err)
	}

	return id, nil
}

func (s *Storage) SelectURL(alias string) (string, error) {
	var resUrl string

	err := s.db.QueryRow("SELECT url FROM url WHERE alias=$1", alias).Scan(&resUrl)
	if err != nil {
		if err == sql.ErrNoRows {
			return "", fmt.Errorf("alias not found")
		}
		return "", fmt.Errorf("SelectURL: get data failed %w", err)
	}

	return resUrl, nil

}
