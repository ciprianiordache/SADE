package server

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func New(addr string, router http.Handler) *Server {
	return &Server{
		httpServer: &http.Server{
			Addr:    addr,
			Handler: router,
		},
	}
}

func (s *Server) start() error {
	fmt.Printf("Server is starting at %s\n", s.httpServer.Addr)

	go func() {
		err := s.httpServer.ListenAndServe()
		if err != nil && err != http.ErrServerClosed {
			fmt.Printf("Error starting server: %s", err)
		}
	}()
	return nil
}

func (s *Server) shutdown() error {
	fmt.Println("Server is shutting down...")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	err := s.httpServer.Shutdown(ctx)
	if err != nil {
		return fmt.Errorf("error shutting down http server: %s", err)
	}
	fmt.Println("Server is off!")
	return nil
}

func RunServer(srv *Server) error {
	err := srv.start()
	if err != nil {
		return err
	}
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, syscall.SIGINT, syscall.SIGTERM)

	select {
	case <-interrupt:
		err := srv.shutdown()
		if err != nil {
			return fmt.Errorf("error shutting down the server: %s", err)
		}
	}

	return nil
}
