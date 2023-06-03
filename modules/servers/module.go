package servers

import "github.com/gofiber/fiber/v2"

type IModule interface {
	NewMonitorModule() IMonitorModule
	NewUsersModule() IUsersModule
	NewProductsModule() IProductsModule
}

type module struct {
	r fiber.Router
	s *server
	m IMiddlewares
}

func NewModule(s *server, m IMiddlewares) IModule {
	return &module{
		r: s.app.Group("v1"),
		s: s,
		m: m,
	}
}
