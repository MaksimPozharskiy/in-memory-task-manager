package requestscounter

import (
	"sync"
	"testing"
)

func TestNewRequestsCounter(t *testing.T) {
	rc := NewRequestsCounter()
	if rc == nil {
		t.Errorf("NewRequestsCounter() should not return nil")
		return
	}
	if len(rc.Counts) != 0 {
		t.Errorf("NewRequestsCounter() should initialize an empty map, got %v", rc.Counts)
	}
}

func TestIncrementAndGet(t *testing.T) {
	rc := NewRequestsCounter()
	id := 1
	rc.Increment(id)
	count := rc.Get(id)
	if count != 1 {
		t.Errorf("Expected count of 1, got %d", count)
	}

	// check repeating increment
	rc.Increment(id)
	count = rc.Get(id)
	if count != 2 {
		t.Errorf("Expected count of 2, got %d", count)
	}
}

func TestIncrementConcurrently(t *testing.T) {
	rc := NewRequestsCounter()
	id := 1
	var wg sync.WaitGroup
	workers := 10

	for i := 0; i < workers; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			rc.Increment(id)
		}()
	}

	wg.Wait()

	if count := rc.Get(id); count != workers {
		t.Errorf("Expected count of %d, got %d", workers, count)
	}
}
