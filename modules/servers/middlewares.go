package servers

import (
	"github.com/Rayato159/neversuitup-e-commerce-test/modules/middlewares/middlewaresHandler"
)

type IMiddlewares interface {
	Handler() middlewaresHandler.IMiddlewaresHandler
}

type middlewares struct {
	*server
	handler middlewaresHandler.IMiddlewaresHandler
}

func (s *server) NewMiddlewares() IMiddlewares {
	handler := middlewaresHandler.MiddlewaresHandler(s.cfg)

	return &middlewares{
		server:  s,
		handler: handler,
	}
}

func (m *middlewares) Handler() middlewaresHandler.IMiddlewaresHandler { return m.handler }
