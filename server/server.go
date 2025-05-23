package server

import (
	"log"
	"net/http"
)

func Start() {
	handler := http.NewServeMux()

	s := &http.Server{
		Addr:         ":8080",
		Handler:      handler,
		ReadTimeout:  10,
		WriteTimeout: 10,
		IdleTimeout:  1 << 20,
	}

	log.Fatal(s.ListenAndServe())
}
