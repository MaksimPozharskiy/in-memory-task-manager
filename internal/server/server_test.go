package server

import (
	"net/http"
	"testing"

	"github.com/MaksimPozharskiy/in-memory-task-manager/internal/requestscounter"
)

func TestNewServer(t *testing.T) {
	handler := http.NewServeMux()
	counter := requestscounter.NewRequestsCounter()
	srv := NewServer(":8080", handler, counter)

	if srv.httpServer.Addr != ":8080" {
		t.Errorf("Expected server address to be ':8080', got '%s'", srv.httpServer.Addr)
	}
	if srv.requestsCounter != counter {
		t.Errorf("Expected requestsCounter to be initialized")
	}
}

func TestActiveRequestsIncrementAndDecrement(t *testing.T) {
	srv := NewServer("", nil, nil)

	srv.IncrementActiveRequests()
	if !srv.IsActiveRequests() {
		t.Errorf("Expected active requests to be incremented")
	}

	srv.DecrementActiveRequests()
	if srv.IsActiveRequests() {
		t.Errorf("Expected active requests to be decremented")
	}
}
