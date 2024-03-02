package api

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"sync"
)


var (
	// map for storing requsets in format `id: count requests`
	requestsCount = make(map[int]int)
	mu            sync.RWMutex
)

func TaskHandler(w http.ResponseWriter, r *http.Request) {
	// get id from parameters
	idStr := strings.TrimPrefix(r.URL.Path, "/api/task/")
	id, err := strconv.Atoi(idStr)
	if err != nil || idStr == "" {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	// comute requests count
	mu.Lock()
	requestsCount[id]++
	count := requestsCount[id]
	mu.Unlock()

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(fmt.Sprintf("%d:%d", id, count)))
}
