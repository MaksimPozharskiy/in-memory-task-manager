package requestscounter

import (
	"sync"
)

type RequestsCounterInterface interface {
	Increment(id int)
	Get(id int) int
}

// struct for storing requests count
type RequestsCounter struct {
	mu     sync.RWMutex
	Counts map[int]int
}

func NewRequestsCounter() *RequestsCounter {
	return &RequestsCounter{
		Counts: make(map[int]int),
	}
}

func (rc *RequestsCounter) Increment(id int) {
	rc.mu.Lock()
	defer rc.mu.Unlock()
	rc.Counts[id]++
}

func (rc *RequestsCounter) Get(id int) int {
	rc.mu.RLock()
	defer rc.mu.RUnlock()
	return rc.Counts[id]
}
