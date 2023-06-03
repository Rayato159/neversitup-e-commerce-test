package servers

import (
	"github.com/Rayato159/neversuitup-e-commerce-test/modules/users/usersHandler"
	"github.com/Rayato159/neversuitup-e-commerce-test/modules/users/usersRepository"
	"github.com/Rayato159/neversuitup-e-commerce-test/modules/users/usersUsecase"
)

type IUsersModule interface {
	Init()
	Handler() usersHandler.IUsersHandler
	Usecase() usersUsecase.IUsersUsecase
	Repository() usersRepository.IUsersRepository
}

type usersModule struct {
	*module
	handler    usersHandler.IUsersHandler
	usecase    usersUsecase.IUsersUsecase
	repository usersRepository.IUsersRepository
}

func (m *module) NewUsersModule() IUsersModule {
	repository := usersRepository.NewUsersRepository(m.s.db)
	usecase := usersUsecase.NewUsersUsercase(repository)
	handler := usersHandler.NewUsersHandler(m.s.cfg, usecase)

	return &usersModule{
		module:     m,
		repository: repository,
		usecase:    usecase,
		handler:    handler,
	}
}

func (m *usersModule) Init() {
	g := m.r.Group("/users")

	g.Get("/:id", m.handler.FindOneUser)
	g.Post("/", m.m.Handler().ApiKeyAuth(), m.handler.Register)
}
func (m *usersModule) Handler() usersHandler.IUsersHandler          { return m.handler }
func (m *usersModule) Usecase() usersUsecase.IUsersUsecase          { return m.usecase }
func (m *usersModule) Repository() usersRepository.IUsersRepository { return m.repository }
