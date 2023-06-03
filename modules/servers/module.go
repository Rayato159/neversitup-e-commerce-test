package servers

import (
	"github.com/gofiber/fiber/v2"
	"google.golang.org/grpc"
)

type IModule interface {
	NewMonitorModule() IMonitorModule
	NewUsersModule() IUsersModule
	NewProductsModule() IProductsModule
	NewAuthModule() IAuthModule
	NewOrdersModule() IOrdersModule
}

type module struct {
	r          fiber.Router
	s          *server
	m          IMiddlewares
	grpcServer *grpc.Server
}

func NewModule(s *server, m IMiddlewares, grpcServer *grpc.Server) IModule {
	return &module{
		r:          s.app.Group("v1"),
		s:          s,
		m:          m,
		grpcServer: grpcServer,
	}
}
