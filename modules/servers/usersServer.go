package servers

import (
	"log"
	"os"
	"os/signal"
)

func (s *server) StartUsersServer() {
	// Middlewares

	// Modules
	modules := NewModule(s, nil)
	modules.NewMonitorModule().Init()

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
