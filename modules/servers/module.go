package servers

import "github.com/gofiber/fiber/v2"

type IModule interface {
	NewMonitorModule() IMonitorModule
}

type module struct {
	r fiber.Router
	s *server
	m any
}

func NewModule(s *server, m any) IModule {
	return &module{
		r: s.App.Group("v1"),
		s: s,
		m: m,
	}
}
