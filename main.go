package main

import (
	"context"
	"net/http"
	"time"

	"github.com/sirupsen/logrus"

	"ClientMicroservice/communications"
	"ClientMicroservice/config"
	"ClientMicroservice/handlers"
	"ClientMicroservice/services"
)

type Server struct {
	httpServer *http.Server
}

func main() {
	logrus.SetFormatter(new(logrus.JSONFormatter))

	cfg, err := config.New()
	if err != nil {
		logrus.Fatalf("failed to initialize config %v", err)
	}

	client := &http.Client{
		Timeout: 60 * time.Second,
	}

	domainService := communications.NewDomain(cfg.DomainServiceURL, client)
	service := services.NewService(domainService)
	handler := handlers.NewHandler(service)

	var srv Server
	if err := srv.Run(cfg.Port, handler.InitRouts()); err != nil {
		_ = srv.Shutdown(context.TODO())
		logrus.Fatalf("error occured while running http server %s", err.Error())
	}
}

func (s *Server) Run(port string, handler http.Handler) error {
	s.httpServer = &http.Server{
		Addr:         ":" + port,
		Handler:      handler,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}
	return s.httpServer.ListenAndServe()
}

func (s *Server) Shutdown(ctx context.Context) error {
	return s.httpServer.Shutdown(ctx)
}
