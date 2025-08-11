package http

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"lms_system/internal/domain"
)

type Server struct {
	httpServer  *http.Server
	service     domain.ServiceInterface
	authService domain.AuthServiceInterface
	fileService domain.FileServiceInterface
}

func NewServer(service domain.ServiceInterface, authService domain.AuthServiceInterface, fileService domain.FileServiceInterface, port string) *Server {
	return &Server{
		service:     service,
		authService: authService,
		fileService: fileService,
		httpServer: &http.Server{
			Addr:         "0.0.0.0:" + port,
			Handler:      NewRouter(service, authService, fileService),
			ReadTimeout:  15 * time.Second,
			WriteTimeout: 15 * time.Second,
			IdleTimeout:  60 * time.Second,
		},
	}
}

//func (s *Server) Start() error {
//	go func() {
//		fmt.Printf("Server starting on %s\n", s.httpServer.Addr)
//		if err := s.httpServer.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
//			fmt.Printf("Server failed to start: %v\n", err)
//		}
//	}()
//
//	return nil
//}

func (s *Server) Start() error {
	fmt.Printf("Server starting on %s\n", s.httpServer.Addr)
	return s.httpServer.ListenAndServe()
}

func (s *Server) Stop() error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	return s.httpServer.Shutdown(ctx)
}
