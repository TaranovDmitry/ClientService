package main

import (
	"context"
	"net/http"
	"os"
	"time"

	"github.com/sirupsen/logrus"

	"gihub.com/TaranovDmitry/ClientService/communications"
	"gihub.com/TaranovDmitry/ClientService/config"
	"gihub.com/TaranovDmitry/ClientService/handlers"
	"gihub.com/TaranovDmitry/ClientService/services"
)

type Server struct {
	httpServer *http.Server
}

func main() {
	logger := logrus.New()
	logger.SetFormatter(new(logrus.JSONFormatter))
	logger.SetOutput(os.Stdout)

	cfg, err := config.New()
	if err != nil {
		logger.Fatalf("failed to initialize config %v", err)
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
		logger.Fatalf("error occured while running http server %s", err.Error())
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
