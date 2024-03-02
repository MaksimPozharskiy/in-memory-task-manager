package api

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/MaksimPozharskiy/in-memory-task-manager/internal/requestscounter"
)

type API struct {
	Counter                 *requestscounter.RequestsCounter
	IncrementActiveRequests func()
	DecrementActiveRequests func()
}

func (api *API) TaskHandler(w http.ResponseWriter, r *http.Request) {
	// compute active requests for graceful shutdown
	api.IncrementActiveRequests()
	defer api.DecrementActiveRequests()

	// get id from parameters
	idStr := strings.TrimPrefix(r.URL.Path, "/api/task/")
	id, err := strconv.Atoi(idStr)
	if err != nil || idStr == "" {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	// increment requests count
	api.Counter.Increment(id)
	count := api.Counter.Get(id)

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(fmt.Sprintf("%d:%d", id, count)))
}
