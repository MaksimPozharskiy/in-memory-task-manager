package main

import (
	"net/http"

	"github.com/MaksimPozharskiy/in-memory-task-manager/internal/api"
	"github.com/MaksimPozharskiy/in-memory-task-manager/internal/requestscounter"
	"github.com/MaksimPozharskiy/in-memory-task-manager/internal/server"
)

func main() {
	requestsCounter := requestscounter.NewRequestsCounter()

	server := server.NewServer(":8081", nil, requestsCounter)
	handler := setupRouter(requestsCounter, server)
	server.SetHandler(handler)

	server.Start()
}

func setupRouter(counter *requestscounter.RequestsCounter, server *server.Server) http.Handler {
	mux := http.NewServeMux()
	apiInstance := &api.API{Counter: counter, IncrementActiveRequests: server.IncrementActiveRequests,
		DecrementActiveRequests: server.DecrementActiveRequests}

	mux.HandleFunc("/api/task/", apiInstance.TaskHandler)

	return mux
}
