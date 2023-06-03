package servers

import (
	"log"
	"os"
	"os/signal"

	"google.golang.org/grpc"
)

func (s *server) StartProductsServer() {
	// Middlewares
	middlewares := s.NewMiddlewares()
	s.app.Use(middlewares.Handler().Logger())
	s.app.Use(middlewares.Handler().Cors())

	var opts []grpc.ServerOption
	grpcServer := grpc.NewServer(opts...)
	defer grpcServer.Stop()

	// Modules
	modules := NewModule(s, middlewares, grpcServer)
	modules.NewMonitorModule().Init()
	modules.NewProductsModule().Init()

	s.app.Use(middlewares.Handler().RouterCheck())

	// Graceful Shutdown
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	go func() {
		_ = <-c
		log.Println("server is shutting down...")
		_ = s.app.Shutdown()
	}()

	// Listen to host:port
	log.Printf("products server is starting on %v", s.cfg.App().Url())
	s.app.Listen(s.cfg.App().Url())
}
