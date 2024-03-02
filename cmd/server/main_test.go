package main

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/MaksimPozharskiy/in-memory-task-manager/internal/requestscounter"
	"github.com/MaksimPozharskiy/in-memory-task-manager/internal/server"
)

func TestSetupRouter(t *testing.T) {
	counter := requestscounter.NewRequestsCounter()
	srv := server.NewServer("", nil, counter)
	handler := setupRouter(counter, srv)

	ts := httptest.NewServer(handler)
	defer ts.Close()

	res, err := http.Get(fmt.Sprintf("%s/api/task/1", ts.URL))
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
	if res.StatusCode != http.StatusOK {
		t.Errorf("Expected status code 200, got %d", res.StatusCode)
	}
}
