package rest

import (
	"context"
	"log"
	"log/slog"
	"net"
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"sync"
	"syscall"

	"github.com/go-chi/chi/v5"
)

type Server struct {
	cfg    *Config
	Router *chi.Mux
	Server *http.Server
	Logger *slog.Logger
}

func NewServer(cfg *Config, logger *slog.Logger, router *chi.Mux) (*Server, error) {
	server := &Server{
		cfg:    cfg,
		Router: router,
		Logger: logger,
		Server: &http.Server{
			ReadTimeout:  cfg.HTTPServerReadTimeout,
			WriteTimeout: cfg.HTTPServerWriteTimeout,
			IdleTimeout:  cfg.HTTPServerIdleTimeout,
			Handler:      router,
		},
	}
	return server, nil
}

// Start the HTTP server and blocks until we either receive a signal or the HTTP server returns an error.
func (s *Server) Start() error {
	var wg sync.WaitGroup
	wg.Add(1)

	// Listen for signals - shutdown the server if we receive one
	go func() {
		quit := make(chan os.Signal, 1)
		signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
		<-quit

		ctx, cancel := context.WithTimeout(context.Background(), s.cfg.ServerGracefulShutdownTimeout)
		defer cancel()

		s.Logger.Info("Stopping HTTP server", "reason", "received signal")
		s.Server.SetKeepAlivesEnabled(false)
		err := s.Server.Shutdown(ctx)
		if err != nil {
			log.Panic(err.Error())
		}

		wg.Done()
	}()

	port := s.cfg.HTTPPort
	listener, err := net.Listen("tcp", net.JoinHostPort(s.cfg.HTTPAddress, strconv.Itoa(port)))
	if err != nil {
		return err
	}
	s.Logger.Info("Server listening on address", "address", listener.Addr().String())

	err = s.Server.Serve(listener)
	if err != http.ErrServerClosed {
		return err
	}

	wg.Wait()
	s.Logger.Info("Stopped HTTP server", "address", listener.Addr().String())

	return nil
}
