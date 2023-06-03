package middlewaresHandler

import (
	"context"
	"strings"
	"time"

	pb "github.com/Rayato159/neversuitup-e-commerce-test/modules/auth/authProto"
	"github.com/Rayato159/neversuitup-e-commerce-test/modules/middlewares"

	"github.com/Rayato159/neversuitup-e-commerce-test/config"
	"github.com/Rayato159/neversuitup-e-commerce-test/modules/auth"
	"github.com/Rayato159/neversuitup-e-commerce-test/modules/entities"
	"github.com/Rayato159/neversuitup-e-commerce-test/pkg/lockkunchae"
	"github.com/Rayato159/neversuitup-e-commerce-test/pkg/utils"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type IMiddlewaresHandler interface {
	Cors() fiber.Handler
	RouterCheck() fiber.Handler
	Logger() fiber.Handler
	JwtAuth() fiber.Handler
	ParamsCheck() fiber.Handler
	ApiKeyAuth() fiber.Handler
}

type middlewaresHandler struct {
	cfg config.IConfig
}

func MiddlewaresHandler(cfg config.IConfig) IMiddlewaresHandler {
	return &middlewaresHandler{
		cfg: cfg,
	}
}

func (h *middlewaresHandler) Cors() fiber.Handler {
	return cors.New(cors.Config{
		Next:             cors.ConfigDefault.Next,
		AllowOrigins:     "*",
		AllowMethods:     "GET,POST,HEAD,PUT,DELETE,PATCH",
		AllowHeaders:     "",
		AllowCredentials: false,
		ExposeHeaders:    "",
		MaxAge:           0,
	})
}

func (h *middlewaresHandler) RouterCheck() fiber.Handler {
	return func(c *fiber.Ctx) error {
		return entities.NewResponse(c).Error(
			fiber.ErrNotFound.Code,
			string(middlewares.RouterCheckErr),
			"rotuer not found",
		).Res()
	}
}

func (h *middlewaresHandler) Logger() fiber.Handler {
	return logger.New(logger.Config{
		Format:     "${time} [${ip}] ${status} - ${method} ${path}\n",
		TimeFormat: "02/01/2006",
		TimeZone:   "Bangkok/Asia",
	})
}

func (h *middlewaresHandler) JwtAuth() fiber.Handler {
	return func(c *fiber.Ctx) error {
		token := strings.TrimPrefix(c.Get("Authorization"), "Bearer ")
		result, err := lockkunchae.ParseToken(h.cfg.Jwt(), token)
		if err != nil {
			return entities.NewResponse(c).Error(
				fiber.ErrUnauthorized.Code,
				string(middlewares.JwtAuthErr),
				err.Error(),
			).Res()
		}

		claims := result.Claims

		conn, err := grpc.Dial(h.cfg.App().AuthGrpcUrl(), grpc.WithTransportCredentials(insecure.NewCredentials()))
		if err != nil {
			return entities.NewResponse(c).Error(
				fiber.ErrInternalServerError.Code,
				string(auth.LoginErr),
				err.Error(),
			).Res()
		}
		defer conn.Close()

		ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
		defer cancel()

		client := pb.NewAuthServicesClient(conn)
		utils.Debug(claims)
		isValid, _ := client.FindAccessToken(ctx, &pb.FindAccessTokenReq{
			UserId:      claims.Id,
			AccessToken: token,
		})
		if isValid == nil {
			return entities.NewResponse(c).Error(
				fiber.ErrInternalServerError.Code,
				string(middlewares.JwtAuthErr),
				"auth grpc server down",
			).Res()
		}
		if !isValid.Result {
			return entities.NewResponse(c).Error(
				fiber.ErrUnauthorized.Code,
				string(middlewares.JwtAuthErr),
				"access_token is invalid",
			).Res()
		}

		// Set UserId
		c.Locals("userId", claims.Id)
		c.Locals("userUsername", claims.Username)
		return c.Next()
	}
}

func (h *middlewaresHandler) ParamsCheck() fiber.Handler {
	return func(c *fiber.Ctx) error {
		userId := c.Locals("userId")
		if c.Params("user_id") != userId {
			return entities.NewResponse(c).Error(
				fiber.ErrUnauthorized.Code,
				string(middlewares.ParamsCheckErr),
				"never gonna give you up",
			).Res()
		}
		return c.Next()
	}
}

func (h *middlewaresHandler) ApiKeyAuth() fiber.Handler {
	return func(c *fiber.Ctx) error {
		key := c.Get("X-Api-Key")
		if _, err := lockkunchae.ParseApiKey(h.cfg.Jwt(), key); err != nil {
			return entities.NewResponse(c).Error(
				fiber.ErrUnauthorized.Code,
				string(middlewares.ApiKeyErr),
				"apikey is invalid or required",
			).Res()
		}
		return c.Next()
	}
}
