package server

import (
	"Service-for-assigning-reviewers-for-Pull-Requests/config"
	"context"
	"fmt"
	"github.com/go-chi/chi/v5"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

const (
	defaultReadTimeout     = 5 * time.Second
	defaultWriteTimeout    = 5 * time.Second
	defaultAddr            = "0.0.0.0:3333"
	defaultShutdownTimeout = 10 * time.Second
)

type Server struct {
	internalServer  *http.Server
	channelErr      chan error
	shutdownTimeout time.Duration
}

func (s *Server) Start() {
	go func() {
		s.channelErr <- s.internalServer.ListenAndServe()
		close(s.channelErr)
	}()
}

func NewServer(handler http.Handler, options ...Option) *Server {
	server := &Server{
		internalServer: &http.Server{
			Handler:      handler,
			ReadTimeout:  defaultReadTimeout,
			WriteTimeout: defaultWriteTimeout,
			Addr:         defaultAddr,
		},
		channelErr:      make(chan error, 1),
		shutdownTimeout: defaultShutdownTimeout,
	}
	for _, option := range options {
		option(server)
	}
	server.Start()
	return server
}

func (s *Server) FullShutdownTimeout(logger *slog.Logger) error {
	ctx, cancel := context.WithTimeout(context.Background(), s.shutdownTimeout)
	defer cancel()
	logger.Info("Shutting down server...\n")
	if err := s.internalServer.Shutdown(ctx); err != nil {
		return fmt.Errorf("server shutdown filed: %v", err)
	}
	return nil
}

func (s *Server) GracefulShutdown(logger *slog.Logger) {
	osInterruptChan := make(chan os.Signal, 1)
	signal.Notify(osInterruptChan, syscall.SIGTERM, syscall.SIGINT)
	timeoutChan := time.After(s.shutdownTimeout)

	select {
	case <-osInterruptChan:
		logger.Info("server interrupted by system or user")
	case <-s.channelErr:
		logger.Error("server error occurred", slog.Any("error", <-s.channelErr))
	case <-timeoutChan:
		logger.Info("shutdown timeout reached")
	}

	ctx, cancel := context.WithTimeout(context.Background(), s.shutdownTimeout)
	defer cancel()

	logger.Info("shutting down server...")
	if err := s.internalServer.Shutdown(ctx); err != nil {
		logger.Error("graceful shutdown failed", slog.Any("error", err))
	} else {
		logger.Info("server stopped successfully")
	}

	close(osInterruptChan)
}

func StartServer(cfg *config.Config, controller *chi.Mux, logger *slog.Logger) *Server {
	customServer := NewServer(controller,
		SetReadTimeout(*cfg.Server.ReadTimeout),
		SetWriteTimeout(*cfg.Server.WriteTimeout),
		SetAddr(cfg.Server.Addr),
		SetShutdownTimeout(cfg.Server.ShutdownTimeout),
	)
	fmt.Println(customServer.internalServer.ReadTimeout,
		customServer.internalServer.WriteTimeout, customServer.internalServer.Addr,
		customServer.shutdownTimeout)
	logger.Info("successfully created server\n")
	return customServer
}
