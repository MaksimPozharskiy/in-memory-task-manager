package api

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/MaksimPozharskiy/in-memory-task-manager/internal/requestscounter"
)

var _ requestscounter.RequestsCounterInterface = &mockRequestsCounter{}

type mockRequestsCounter struct {
	counts map[int]int
}

func newMockRequestsCounter() *mockRequestsCounter {
	return &mockRequestsCounter{counts: make(map[int]int)}
}

func (m *mockRequestsCounter) Increment(id int) {
	m.counts[id]++
}

func (m *mockRequestsCounter) Get(id int) int {
	return m.counts[id]
}

func TestTaskHandler_ValidID(t *testing.T) {
	counter := newMockRequestsCounter()

	api := API{
		Counter:                 counter,
		IncrementActiveRequests: func() {},
		DecrementActiveRequests: func() {},
	}

	req, _ := http.NewRequest("GET", "/api/task/1", nil)
	rr := httptest.NewRecorder()

	handler := http.HandlerFunc(api.TaskHandler)
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	expected := "1:1"
	if rr.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v", rr.Body.String(), expected)
	}
}

func TestTaskHandler_InvalidID(t *testing.T) {
	counter := newMockRequestsCounter()
	api := API{
		Counter:                 counter,
		IncrementActiveRequests: func() {},
		DecrementActiveRequests: func() {},
	}

	req, _ := http.NewRequest("GET", "/api/task/invalid", nil)
	rr := httptest.NewRecorder()

	handler := http.HandlerFunc(api.TaskHandler)
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusBadRequest {
		t.Errorf("handler returned wrong status code for invalid ID: got %v want %v", status, http.StatusBadRequest)
	}
}
