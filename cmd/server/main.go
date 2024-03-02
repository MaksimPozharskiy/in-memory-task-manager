package main

import (
	"log"
	"net/http"
	"time"

	"github.com/MaksimPozharskiy/in-memory-task-manager/internal/api"
	"github.com/MaksimPozharskiy/in-memory-task-manager/internal/requestscounter"
	"github.com/MaksimPozharskiy/in-memory-task-manager/internal/shutdown"
)

func main() {
	requestsCounter := requestscounter.NewRequestsCounter()

	server := &http.Server{
		Addr:         ":8081",
		Handler:      setupRouter(requestsCounter),
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

func setupRouter(counter *requestscounter.RequestsCounter) http.Handler {
	mux := http.NewServeMux()
	apiInstance := &api.API{Counter: counter}

	mux.HandleFunc("/api/task/", apiInstance.TaskHandler)

	return mux
}
