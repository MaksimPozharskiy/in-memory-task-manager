package shutdown

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/MaksimPozharskiy/in-memory-task-manager/internal/requestscounter"
)

func GracefulShutdown(server *http.Server, requestsCounter *requestscounter.RequestsCounter) {
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGINT, syscall.SIGTERM)

	<-stop

	log.Println("Shutting down server...")

	if requestsCounter.Counts != nil {
		log.Println("Requests statistics:")
		for id, count := range requestsCounter.Counts {
			log.Printf("ID %d: %d requests\n", id, count)
		}
	}

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		log.Fatalf("Graceful shutdown failed: %s\n", err)
	} else {
		log.Println("Server gracefully stopped")
	}
}
