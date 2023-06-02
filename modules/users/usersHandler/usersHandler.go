package usersHandler

import (
	"strings"

	"github.com/Rayato159/neversuitup-e-commerce-test/config"
	"github.com/Rayato159/neversuitup-e-commerce-test/modules/entities"
	"github.com/Rayato159/neversuitup-e-commerce-test/modules/users"
	"github.com/Rayato159/neversuitup-e-commerce-test/modules/users/usersUsecase"
	"github.com/gofiber/fiber/v2"
)

type IUsersHandler interface {
	Register(c *fiber.Ctx) error
	FindOneUser(c *fiber.Ctx) error
}

type usersHandler struct {
	cfg          config.IConfig
	usersUsecase usersUsecase.IUsersUsecase
}

func NewUsersHandler(cfg config.IConfig, usersUsecase usersUsecase.IUsersUsecase) IUsersHandler {
	return &usersHandler{
		cfg:          cfg,
		usersUsecase: usersUsecase,
	}
}

func (h *usersHandler) Register(c *fiber.Ctx) error {
	req := new(users.UserCredential)
	if err := c.BodyParser(req); err != nil {
		return entities.NewResponse(c).Error(
			fiber.ErrBadRequest.Code,
			string(users.RegisterErr),
			err.Error(),
		).Res()
	}

	user, err := h.usersUsecase.InsertUser(req)
	if err != nil {
		return entities.NewResponse(c).Error(
			fiber.ErrBadRequest.Code,
			string(users.RegisterErr),
			err.Error(),
		).Res()
	}

	return entities.NewResponse(c).Success(
		fiber.StatusCreated,
		user,
	).Res()
}

func (h *usersHandler) FindOneUser(c *fiber.Ctx) error {
	userId := strings.Trim(c.Params("id"), " ")

	user, err := h.usersUsecase.FindOneUser(userId)
	if err != nil {
		return entities.NewResponse(c).Error(
			fiber.ErrBadRequest.Code,
			string(users.RegisterErr),
			err.Error(),
		).Res()
	}

	return entities.NewResponse(c).Success(
		fiber.StatusCreated,
		user,
	).Res()
}
