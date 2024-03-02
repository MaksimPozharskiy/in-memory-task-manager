package server

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"sync/atomic"
	"syscall"
	"time"

	"github.com/MaksimPozharskiy/in-memory-task-manager/internal/requestscounter"
)

type Server struct {
	httpServer      *http.Server
	shutdownSignal  chan os.Signal
	requestsCounter *requestscounter.RequestsCounter
	activeRequests  int32
}

func NewServer(addr string, handler http.Handler, counter *requestscounter.RequestsCounter) *Server {
	return &Server{
		httpServer: &http.Server{
			Addr:    addr,
			Handler: handler,
		},
		requestsCounter: counter,
		shutdownSignal:  make(chan os.Signal, 1),
	}
}

func (s *Server) Start() {
	signal.Notify(s.shutdownSignal, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		log.Printf("Server is starting on %s\n", s.httpServer.Addr)
		if err := s.httpServer.ListenAndServe(); err != http.ErrServerClosed {
			log.Fatalf("HTTP server ListenAndServe: %v", err)
		}
	}()

	<-s.shutdownSignal
	log.Println("Shutdown signal received, initiating graceful shutdown...")
	s.GracefulShutdown()
}

func (s *Server) GracefulShutdown() {
	log.Println("Shutting down server...")

	// check active connections
	for s.IsActiveRequests() {
		log.Println("Waiting for active requests to complete...")
		// @TODO improve with  addition chan??
		time.Sleep(1 * time.Second)
	}

	// printing requests history
	if s.requestsCounter != nil {
		log.Println("Requests statistics:")
		for id, count := range s.requestsCounter.Counts {
			log.Printf("ID %d: %d requests\n", id, count)
		}
	}

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	if err := s.httpServer.Shutdown(ctx); err != nil {
		log.Fatalf("Graceful shutdown failed: %s\n", err)
	} else {
		log.Println("Server gracefully stopped")
	}
}

func (s *Server) IncrementActiveRequests() {
	atomic.AddInt32(&s.activeRequests, 1)
}

func (s *Server) DecrementActiveRequests() {
	atomic.AddInt32(&s.activeRequests, -1)
}

func (s *Server) IsActiveRequests() bool {
	return atomic.LoadInt32(&s.activeRequests) > 0
}

func (s *Server) SetHandler(handler http.Handler) {
	s.httpServer.Handler = handler
}
