package main

import (
	"log"
	"net/http"

	"github.com/abhaysoni512/students-api/internal/config"
)

func main() {
	// load config
	cfg := config.Mustload()
	// database setup
	// setup router
	router := http.NewServeMux()
	router.HandleFunc("GET /", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("welcome to student's api"))
	})
	// setup server
	server := http.Server{
		Addr:    cfg.Address,
		Handler: router,
	}
	log.Println("Server started")
	err := server.ListenAndServe()

	if err != nil {
		log.Fatal("failed to start server")
	}

}
