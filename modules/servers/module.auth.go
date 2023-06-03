package servers

import (
	"github.com/Rayato159/neversuitup-e-commerce-test/modules/auth/authHandler"
	"github.com/Rayato159/neversuitup-e-commerce-test/modules/auth/authRepository"
	"github.com/Rayato159/neversuitup-e-commerce-test/modules/auth/authUsecase"
)

type IAuthModule interface {
	Init()
	Handler() authHandler.IAuthHandler
	Usecase() authUsecase.IAuthUsecase
	Repository() authRepository.IAuthRepository
}

type authModule struct {
	*module
	handler    authHandler.IAuthHandler
	usecase    authUsecase.IAuthUsecase
	repository authRepository.IAuthRepository
}

func (m *module) NewAuthModule() IAuthModule {
	repository := authRepository.NewAuthRepository(m.s.db)
	usecase := authUsecase.NewAuthUsecase(repository)
	handler := authHandler.NewAuthHandler(m.s.cfg, usecase)

	return &authModule{
		module:     m,
		repository: repository,
		usecase:    usecase,
		handler:    handler,
	}
}

func (m *authModule) Init() {
	g := m.r.Group("/auth")

	g.Post("/login", m.m.Handler().ApiKeyAuth(), m.handler.Login)
	g.Post("/refresh", m.m.Handler().ApiKeyAuth(), m.handler.Refresh)
}
func (m *authModule) Handler() authHandler.IAuthHandler          { return m.handler }
func (m *authModule) Usecase() authUsecase.IAuthUsecase          { return m.usecase }
func (m *authModule) Repository() authRepository.IAuthRepository { return m.repository }
