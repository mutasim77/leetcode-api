package main

import (
	"log"
	"net/http"

	"github.com/mutasim77/leetcode-api/internal/api"
	"github.com/mutasim77/leetcode-api/pkg/config"
)

func main() {
	cfg := config.Load()

	http.HandleFunc("/user/", api.GetUserHandler)

	log.Printf("Server starting on port %s...\n", cfg.Port)
	log.Fatal(http.ListenAndServe(":"+cfg.Port, nil))
}
