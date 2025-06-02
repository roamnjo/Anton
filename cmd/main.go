package main

import (
	"log"

	"github.com/joho/godotenv"
	"github.com/roamnjo/Anton/server"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Println("Unable load env:", err)
	}
}

func main() {
	server.Run()
}
