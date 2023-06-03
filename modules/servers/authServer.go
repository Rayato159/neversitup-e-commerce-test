package servers

import (
	"log"
	"os"
	"os/signal"
)

func (s *server) StartAuthServer() {
	// Middlewares
	middlewares := s.NewMiddlewares()
	s.app.Use(middlewares.Handler().Logger())
	s.app.Use(middlewares.Handler().Cors())

	// Modules
	modules := NewModule(s, nil)
	modules.NewMonitorModule().Init()

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
