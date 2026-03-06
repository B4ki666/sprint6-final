package server

import (
	"log"
	"net/http"
	"time"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/internal/handlers"
)

// MyServer server structure
type MyServer struct {
	Logger     *log.Logger
	HTTPServer *http.Server
}

// NewServer creates a router, accepts *log.Logger and returns the server structure
func NewServer(logger *log.Logger) *MyServer {
	mux := http.NewServeMux()

	mux.HandleFunc("/", handlers.IndexHandler)
	mux.HandleFunc("/upload", handlers.UploadHandler)

	srv := &http.Server{
		Addr:         ":8080",
		Handler:      mux,
		ErrorLog:     logger,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  15 * time.Second,
	}

	return &MyServer{
		Logger:     logger,
		HTTPServer: srv,
	}
}

// Start starts the server
func (s *MyServer) Start() error {
	return s.HTTPServer.ListenAndServe()
}
