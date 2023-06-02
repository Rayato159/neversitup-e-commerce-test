package usersServer

import (
	"log"
	"os"
	"os/signal"

	"github.com/Rayato159/neversuitup-e-commerce-test/modules/servers"
)

type IUsersServer interface {
	Start()
}

type usersServer struct {
	*servers.Server
}

func NewUsersServer(s *servers.Server) *usersServer {
	return &usersServer{s}
}

func (s *usersServer) Start() {
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
	log.Printf("users server is starting on %v", s.Cfg.App().Url())
	s.App.Listen(s.Cfg.App().Url())
}
