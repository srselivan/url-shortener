package server

import (
	"context"
	"log"
	"net"
	"net/http"
	"time"
	"url-shortener/config"
)

const (
	defaultWriteTimeout = 15 * time.Second
	defaultReadTimeout  = 15 * time.Second
	defaultIdleTimeout  = 60 * time.Second
)

type Server struct {
	httpServer *http.Server
}

func New(handler http.Handler, cfg *config.Config) *Server {
	return &Server{
		httpServer: &http.Server{
			Addr:         net.JoinHostPort("", cfg.HTTP.Port),
			WriteTimeout: defaultWriteTimeout,
			ReadTimeout:  defaultReadTimeout,
			IdleTimeout:  defaultIdleTimeout,
			Handler:      handler,
		},
	}
}

func (s *Server) Start() error {
	log.Printf("starting server on %s port", s.httpServer.Addr)
	return s.httpServer.ListenAndServe()
}

func (s *Server) Stop(ctx context.Context) error {
	log.Println("stopping server")
	return s.httpServer.Shutdown(ctx)
}
