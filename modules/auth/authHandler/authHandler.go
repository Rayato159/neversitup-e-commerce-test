package authHandler

import (
	"context"
	"time"

	"github.com/Rayato159/neversuitup-e-commerce-test/config"
	"github.com/Rayato159/neversuitup-e-commerce-test/modules/auth"
	"github.com/Rayato159/neversuitup-e-commerce-test/modules/auth/authUsecase"
	"github.com/Rayato159/neversuitup-e-commerce-test/modules/entities"
	"github.com/Rayato159/neversuitup-e-commerce-test/modules/users"
	pb "github.com/Rayato159/neversuitup-e-commerce-test/modules/users/usersProto"
	"github.com/gofiber/fiber/v2"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type IAuthHandler interface {
	Login(c *fiber.Ctx) error
}

type authHandler struct {
	cfg         config.IConfig
	authUsecase authUsecase.IAuthUsecase
}

func NewAuthHandler(cfg config.IConfig, authUsecase authUsecase.IAuthUsecase) IAuthHandler {
	return &authHandler{
		cfg:         cfg,
		authUsecase: authUsecase,
	}
}

func (u *authHandler) Login(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*30)
	defer cancel()

	conn, err := grpc.Dial("localhost:1235", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return entities.NewResponse(c).Error(
			fiber.ErrInternalServerError.Code,
			string(auth.LoginErr),
			err.Error(),
		).Res()
	}
	defer conn.Close()

	client := pb.NewUsersServicesClient(conn)

	req := new(users.UserCredential)
	if err := c.BodyParser(req); err != nil {
		return entities.NewResponse(c).Error(
			fiber.ErrBadRequest.Code,
			string(auth.LoginErr),
			err.Error(),
		).Res()
	}

	r, err := client.FindOneUserByUsername(ctx, &pb.FindOneUserByUsernameReq{
		Username: req.Username,
	})
	if err != nil {
		return entities.NewResponse(c).Error(
			fiber.ErrInternalServerError.Code,
			string(auth.LoginErr),
			err.Error(),
		).Res()
	}

	return entities.NewResponse(c).Success(
		fiber.StatusOK,
		&users.UserForAll{
			Id:       r.Id,
			Username: r.Username,
			Password: r.Password,
		},
	).Res()
}
