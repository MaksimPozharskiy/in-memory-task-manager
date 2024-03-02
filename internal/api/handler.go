package api

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"
)

func TaskHandler(w http.ResponseWriter, r *http.Request) {
	idStr := strings.TrimPrefix(r.URL.Path, "/api/task/")
	id, err := strconv.Atoi(idStr)
	if err != nil || idStr == "" {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(fmt.Sprintf("Task ID: %d", id)))
}
