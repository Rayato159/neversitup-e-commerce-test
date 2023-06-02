package servers

import "github.com/gofiber/fiber/v2"

type IModule interface {
	NewMonitorModule() IMonitorModule
	NewUsersModule() IUsersModule
}

type module struct {
	r fiber.Router
	s *server
	m any
}

func NewModule(s *server, m any) IModule {
	return &module{
		r: s.app.Group("v1"),
		s: s,
		m: m,
	}
}
