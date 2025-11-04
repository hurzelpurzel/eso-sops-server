package server

import (
	"context"
	"log"
	"net/http"
	"time"
)

// Server wraps http.Server with helpers
type Server struct {
	srv *http.Server
}

// NewServer creates a new Server instance
func NewServer(addr string, handler http.Handler) *Server {
	return &Server{
		srv: &http.Server{
			Addr:              addr,
			Handler:           handler,
			ReadHeaderTimeout: 5 * time.Second,
			ReadTimeout:       10 * time.Second,
			WriteTimeout:      10 * time.Second,
			IdleTimeout:       120 * time.Second,
		},
	}
}

// ListenAndServe runs the server until ctx is cancelled or an error occurs
func (s *Server) ListenAndServe(ctx context.Context) error {
	// start listen in background
	errc := make(chan error, 1)
	go func() {
		log.Printf("http server listening on %s", s.srv.Addr)
		errc <- s.srv.ListenAndServe()
	}()

	select {
	case <-ctx.Done():
		// graceful shutdown with timeout
		shutdownCtx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()
		return s.srv.Shutdown(shutdownCtx)
	case err := <-errc:
		return err
	}
}
