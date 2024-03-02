package main

import (
	"log"
	"net/http"
	"time"

	"github.com/MaksimPozharskiy/in-memory-task-manager/internal/api"
	"github.com/MaksimPozharskiy/in-memory-task-manager/internal/shutdown"
)

func main() {
	server := &http.Server{
		Addr:         ":8081",
		Handler:      setupRouter(),
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	go func() {
		log.Printf("Starting server on %s\n", server.Addr)
		if err := server.ListenAndServe(); err != http.ErrServerClosed {
			log.Fatalf("Error starting server: %s\n", err)
		}
	}()

	shutdown.GracefulShutdown(server)
}

func setupRouter() http.Handler {
	mux := http.NewServeMux()

	mux.HandleFunc("/api/task/", api.TaskHandler)

	return mux
}
