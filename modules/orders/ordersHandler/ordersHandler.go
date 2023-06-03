package ordersHandler

import (
	"strings"

	"github.com/Rayato159/neversuitup-e-commerce-test/config"
	"github.com/Rayato159/neversuitup-e-commerce-test/modules/entities"
	"github.com/Rayato159/neversuitup-e-commerce-test/modules/orders"
	"github.com/Rayato159/neversuitup-e-commerce-test/modules/orders/ordersUsecase"
	"github.com/gofiber/fiber/v2"
)

type IOrdersHandler interface {
	FindOrders(c *fiber.Ctx) error
	FindOneOrder(c *fiber.Ctx) error
}

type ordersHandler struct {
	cfg           config.IConfig
	ordersUsecase ordersUsecase.IOrdersUsecase
}

func NewOrdersHandler(cfg config.IConfig, ordersUsecase ordersUsecase.IOrdersUsecase) IOrdersHandler {
	return &ordersHandler{
		cfg:           cfg,
		ordersUsecase: ordersUsecase,
	}
}

func (h *ordersHandler) FindOrders(c *fiber.Ctx) error {
	userId := strings.Trim(c.Params("user_id"), " ")

	return entities.NewResponse(c).Success(
		fiber.StatusOK,
		h.ordersUsecase.FindOrders(userId),
	).Res()
}

func (h *ordersHandler) FindOneOrder(c *fiber.Ctx) error {
	userId := strings.Trim(c.Params("user_id"), " ")
	orderId := strings.Trim(c.Params("order_id"), " ")

	order, err := h.ordersUsecase.FindOneOrder(userId, orderId)
	if err != nil {
		return entities.NewResponse(c).Error(
			fiber.ErrBadRequest.Code,
			string(orders.FindOneOrderErr),
			err.Error(),
		).Res()
	}

	return entities.NewResponse(c).Success(
		fiber.StatusOK,
		order,
	).Res()
}
