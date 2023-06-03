package servers

import (
	"github.com/Rayato159/neversuitup-e-commerce-test/modules/orders/ordersHandler"
	"github.com/Rayato159/neversuitup-e-commerce-test/modules/orders/ordersRepository"
	"github.com/Rayato159/neversuitup-e-commerce-test/modules/orders/ordersUsecase"
)

type IOrdersModule interface {
	Init()
	Handler() ordersHandler.IOrdersHandler
	Usecase() ordersUsecase.IOrdersUsecase
	Repository() ordersRepository.IOrdersRepository
}

type ordersModule struct {
	*module
	handler    ordersHandler.IOrdersHandler
	usecase    ordersUsecase.IOrdersUsecase
	repository ordersRepository.IOrdersRepository
}

func (m *module) NewOrdersModule() IOrdersModule {
	repository := ordersRepository.NewOrdersRepository(m.s.db)
	usecase := ordersUsecase.NewOrdersUsecase(repository)
	handler := ordersHandler.NewOrdersHandler(m.s.cfg, usecase)

	return &ordersModule{
		module:     m,
		repository: repository,
		usecase:    usecase,
		handler:    handler,
	}
}

func (m *ordersModule) Init() {
	g := m.r.Group("/orders")

	g.Get("/", m.m.Handler().JwtAuth(), m.handler.FindOrders)
	g.Get("/:order_id", m.m.Handler().JwtAuth(), m.handler.FindOneOrder)
	g.Patch("/:order_id/cancel", m.m.Handler().JwtAuth(), m.handler.CancelOrder)
}
func (m *ordersModule) Handler() ordersHandler.IOrdersHandler          { return m.handler }
func (m *ordersModule) Usecase() ordersUsecase.IOrdersUsecase          { return m.usecase }
func (m *ordersModule) Repository() ordersRepository.IOrdersRepository { return m.repository }
