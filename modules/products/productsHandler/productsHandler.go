package productsHandler

import (
	"github.com/Rayato159/neversuitup-e-commerce-test/config"
	"github.com/Rayato159/neversuitup-e-commerce-test/modules/products/productsUsecase"
)

type IProductsHandler interface{}

type productsHandler struct {
	cfg config.IConfig
	productsUsecase productsUsecase.IProductsUsecase
}

func NewProductsHandler(cfg config.IConfig, productsUsecase productsUsecase.IProductsUsecase) IProductsHandler {
	return &productsHandler{
		cfg: cfg,
		productsUsecase: productsUsecase,
	}
}
