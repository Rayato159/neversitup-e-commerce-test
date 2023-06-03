package servers

import (
	"github.com/Rayato159/neversuitup-e-commerce-test/modules/products/productsHandler"
	"github.com/Rayato159/neversuitup-e-commerce-test/modules/products/productsRepository"
	"github.com/Rayato159/neversuitup-e-commerce-test/modules/products/productsUsecase"
)

type IProductsModule interface {
	Init()
	Handler() productsHandler.IProductsHandler
	Usecase() productsUsecase.IProductsUsecase
	Repository() productsRepository.IProductsRepository
}

type productsModule struct {
	*module
	handler    productsHandler.IProductsHandler
	usecase    productsUsecase.IProductsUsecase
	repository productsRepository.IProductsRepository
}

func (m *module) NewProductsModule() IProductsModule {
	repository := productsRepository.NewProductsRepository(m.s.db)
	usecase := productsUsecase.NewProductsUsecase(repository)
	handler := productsHandler.NewProductsHandler(m.s.cfg, usecase)

	return &productsModule{
		module:     m,
		repository: repository,
		usecase:    usecase,
		handler:    handler,
	}
}

func (m *productsModule) Init() {
	g := m.r.Group("/products")
	_ = g
}
func (m *productsModule) Handler() productsHandler.IProductsHandler          { return m.handler }
func (m *productsModule) Usecase() productsUsecase.IProductsUsecase          { return m.usecase }
func (m *productsModule) Repository() productsRepository.IProductsRepository { return m.repository }
