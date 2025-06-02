package server

import (
	"log"
	"net/http"
	"time"

	"github.com/roamnjo/Anton/data"
	"github.com/roamnjo/Anton/handlers"
)

func Run() {
	r := http.NewServeMux()

	storage, err := data.ConnectDB()
	if err != nil {
		log.Fatal("Error connecting db:", err)
	}

	handler := handlers.NewHandler(storage)

	r.HandleFunc("/save_URL", handler.PostUrl)
	r.HandleFunc("/get", handler.GetUrl)

	s := &http.Server{
		Addr:         ":8080",
		Handler:      r,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  1 << 20,
	}

	log.Println("Server starting on port 8080")
	log.Fatal(s.ListenAndServe())
}
