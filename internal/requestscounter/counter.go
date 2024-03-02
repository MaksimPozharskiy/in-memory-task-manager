package requestscounter

import (
	"sync"
)

// struct for storing requests count
type RequestsCounter struct {
	mu     sync.RWMutex
	counts map[int]int
}

func NewRequestsCounter() *RequestsCounter {
	return &RequestsCounter{
		counts: make(map[int]int),
	}
}

func (rc *RequestsCounter) Increment(id int) {
	rc.mu.Lock()
	defer rc.mu.Unlock()
	rc.counts[id]++
}

func (rc *RequestsCounter) Get(id int) int {
	rc.mu.RLock()
	defer rc.mu.RUnlock()
	return rc.counts[id]
}
