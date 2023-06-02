package productsServer

import (
	"log"
	"os"
	"os/signal"

	"github.com/Rayato159/neversuitup-e-commerce-test/modules/servers"
)

type IProductsServer interface {
	Start()
}

type productsServer struct {
	*servers.Server
}

func NewProductsServer(s *servers.Server) *productsServer {
	return &productsServer{s}
}

func (s *productsServer) Start() {
	// Middlewares

	// Modules
	v1 := s.App.Group("v1")
	_ = v1

	// Graceful Shutdown
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	go func() {
		_ = <-c
		log.Println("server is shutting down...")
		_ = s.App.Shutdown()
	}()

	// Listen to host:port
	log.Printf("products server is starting on %v", s.Cfg.App().Url())
	s.App.Listen(s.Cfg.App().Url())
}
