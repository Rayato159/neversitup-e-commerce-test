package servers

import (
	"github.com/Rayato159/neversuitup-e-commerce-test/modules/middlewares/middlewaresHandler"
	"github.com/Rayato159/neversuitup-e-commerce-test/modules/middlewares/middlewaresRepository"
	"github.com/Rayato159/neversuitup-e-commerce-test/modules/middlewares/middlewaresUsecase"
)

type IMiddlewares interface {
	Repository() middlewaresRepository.IMiddlewaresRepository
	Usecase() middlewaresUsecase.IMiddlewaresUsecase
	Handler() middlewaresHandler.IMiddlewaresHandler
}

type middlewares struct {
	*server
	repository middlewaresRepository.IMiddlewaresRepository
	usecase    middlewaresUsecase.IMiddlewaresUsecase
	handler    middlewaresHandler.IMiddlewaresHandler
}

func (s *server) NewMiddlewares() IMiddlewares {
	repository := middlewaresRepository.MiddlewaresRepository(s.db)
	usecase := middlewaresUsecase.MiddlewaresUsecase(repository)
	handler := middlewaresHandler.MiddlewaresHandler(s.cfg, usecase)

	return &middlewares{
		server:     s,
		repository: repository,
		usecase:    usecase,
		handler:    handler,
	}
}

func (m *middlewares) Repository() middlewaresRepository.IMiddlewaresRepository {
	return m.repository
}
func (m *middlewares) Usecase() middlewaresUsecase.IMiddlewaresUsecase { return m.usecase }
func (m *middlewares) Handler() middlewaresHandler.IMiddlewaresHandler { return m.handler }
