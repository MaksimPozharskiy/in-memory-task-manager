package api

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/MaksimPozharskiy/in-memory-task-manager/internal/requestscounter"
)

type API struct {
	Counter *requestscounter.RequestsCounter
}

func (api *API) TaskHandler(w http.ResponseWriter, r *http.Request) {
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
