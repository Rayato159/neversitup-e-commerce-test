package productsHandler

import (
	"strings"

	"github.com/Rayato159/neversuitup-e-commerce-test/config"
	"github.com/Rayato159/neversuitup-e-commerce-test/modules/entities"
	"github.com/Rayato159/neversuitup-e-commerce-test/modules/products"
	"github.com/Rayato159/neversuitup-e-commerce-test/modules/products/productsUsecase"
	"github.com/gofiber/fiber/v2"
)

type IProductsHandler interface {
	FindProducts(c *fiber.Ctx) error
	FindOneProduct(c *fiber.Ctx) error
}

type productsHandler struct {
	cfg             config.IConfig
	productsUsecase productsUsecase.IProductsUsecase
}

func NewProductsHandler(cfg config.IConfig, productsUsecase productsUsecase.IProductsUsecase) IProductsHandler {
	return &productsHandler{
		cfg:             cfg,
		productsUsecase: productsUsecase,
	}
}

func (h *productsHandler) FindProducts(c *fiber.Ctx) error {
	return entities.NewResponse(c).Success(
		fiber.StatusOK,
		h.productsUsecase.FindProducts(),
	).Res()
}

func (h *productsHandler) FindOneProduct(c *fiber.Ctx) error {
	productId := strings.Trim(c.Params("product_id"), " ")

	product, err := h.productsUsecase.FindOneProduct(productId)
	if err != nil {
		return entities.NewResponse(c).Error(
			fiber.ErrBadRequest.Code,
			string(products.FindOneProductErr),
			err.Error(),
		).Res()
	}

	return entities.NewResponse(c).Success(
		fiber.StatusOK,
		product,
	).Res()
}
