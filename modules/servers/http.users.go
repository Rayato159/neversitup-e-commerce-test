package servers

import (
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"

	pb "github.com/Rayato159/neversuitup-e-commerce-test/modules/users/usersProto"

	"google.golang.org/grpc"
)

func (s *server) StartUsersServer() {
	// Middlewares
	middlewares := s.NewMiddlewares()
	s.app.Use(middlewares.Handler().Logger())
	s.app.Use(middlewares.Handler().Cors())

	// Modules
	modules := NewModule(s, middlewares, nil)
	modules.NewMonitorModule().Init()
	modules.NewUsersModule().Init()

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
	log.Printf("users server is starting on %v", s.cfg.App().Url())
	s.app.Listen(s.cfg.App().Url())
}

func (s *server) StartUsersGRPCServer() {
	var opts []grpc.ServerOption
	grpcServer := grpc.NewServer(opts...)
	lis, err := net.Listen("tcp", fmt.Sprintf("%s:%d", s.cfg.App().Host(), s.cfg.App().Port()))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	modules := NewModule(s, nil, grpcServer)
	pb.RegisterUsersServicesServer(grpcServer, NewUsersGRPCModule(modules.NewUsersModule().Usecase()))

	log.Printf("users grpc server is starting on %v", s.cfg.App().Url())
	grpcServer.Serve(lis)
}
