package main

import (
	"log"
	"net/http"

	"your_project/config"
	"your_project/routes"
)

func main() {
	cfg := config.Load()

	router := routes.RegisterRoutes()

	log.Println("Server running on port", cfg.Port)
	log.Fatal(http.ListenAndServe(":"+cfg.Port, router))
}
