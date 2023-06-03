package ordersHandler

import (
	"context"
	"fmt"
	"log"
	"strings"
	"time"

	pbProducts "github.com/Rayato159/neversuitup-e-commerce-test/modules/products/productsProto"
	pbUsers "github.com/Rayato159/neversuitup-e-commerce-test/modules/users/usersProto"

	"github.com/Rayato159/neversuitup-e-commerce-test/config"
	"github.com/Rayato159/neversuitup-e-commerce-test/modules/auth"
	"github.com/Rayato159/neversuitup-e-commerce-test/modules/entities"
	"github.com/Rayato159/neversuitup-e-commerce-test/modules/orders"
	"github.com/Rayato159/neversuitup-e-commerce-test/modules/orders/ordersUsecase"
	"github.com/gofiber/fiber/v2"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type IOrdersHandler interface {
	FindOrders(c *fiber.Ctx) error
	FindOneOrder(c *fiber.Ctx) error
	CancelOrder(c *fiber.Ctx) error
	InsertOrder(c *fiber.Ctx) error
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
	userId := strings.Trim(c.Locals("userId").(string), " ")
	if userId == "" {
		return entities.NewResponse(c).Error(
			fiber.ErrBadRequest.Code,
			string(orders.FindOrdersErr),
			"user_id is required",
		).Res()
	}

	return entities.NewResponse(c).Success(
		fiber.StatusOK,
		h.ordersUsecase.FindOrders(userId),
	).Res()
}

func (h *ordersHandler) FindOneOrder(c *fiber.Ctx) error {
	userId := strings.Trim(c.Locals("userId").(string), " ")
	if userId == "" {
		return entities.NewResponse(c).Error(
			fiber.ErrBadRequest.Code,
			string(orders.FindOneOrderErr),
			"user_id is required",
		).Res()
	}

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

func (h *ordersHandler) CancelOrder(c *fiber.Ctx) error {
	userId := strings.Trim(c.Locals("userId").(string), " ")
	if userId == "" {
		return entities.NewResponse(c).Error(
			fiber.ErrBadRequest.Code,
			string(orders.CancelOrderErr),
			"user_id is required",
		).Res()
	}

	orderId := strings.Trim(c.Params("order_id"), " ")

	order, err := h.ordersUsecase.CancelOrder(userId, orderId)
	if err != nil {
		return entities.NewResponse(c).Error(
			fiber.ErrBadRequest.Code,
			string(orders.CancelOrderErr),
			err.Error(),
		).Res()
	}

	return entities.NewResponse(c).Success(
		fiber.StatusOK,
		order,
	).Res()
}

func (h *ordersHandler) InsertOrder(c *fiber.Ctx) error {
	userId := strings.Trim(c.Locals("userId").(string), " ")
	if userId == "" {
		return entities.NewResponse(c).Error(
			fiber.ErrBadRequest.Code,
			string(orders.InsertOrderErr),
			"user_id is required",
		).Res()
	}

	req := &orders.Order{
		Products: make([]*orders.OrderProduct, 0),
	}
	if err := c.BodyParser(req); err != nil {
		return entities.NewResponse(c).Error(
			fiber.ErrBadRequest.Code,
			string(orders.InsertOrderErr),
			err.Error(),
		).Res()
	}
	if len(req.Products) == 0 {
		return entities.NewResponse(c).Error(
			fiber.ErrBadRequest.Code,
			string(orders.InsertOrderErr),
			"products is required",
		).Res()
	}
	if req.Address == "" {
		return entities.NewResponse(c).Error(
			fiber.ErrBadRequest.Code,
			string(orders.InsertOrderErr),
			"address is required",
		).Res()
	}
	if req.Contact == "" {
		return entities.NewResponse(c).Error(
			fiber.ErrBadRequest.Code,
			string(orders.InsertOrderErr),
			"contact is required",
		).Res()
	}

	req.UserId = userId

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*30)
	defer cancel()

	// Check user is exists
	connUser, err := grpc.Dial(h.cfg.App().UsersGrpcUrl(), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return entities.NewResponse(c).Error(
			fiber.ErrInternalServerError.Code,
			string(orders.InsertOrderErr),
			err.Error(),
		).Res()
	}
	defer connUser.Close()

	clientUser := pbUsers.NewUsersServicesClient(connUser)

	user, _ := clientUser.FindOneUserById(ctx, &pbUsers.FindOneUserByIdReq{Id: userId})
	if !user.Result {
		return entities.NewResponse(c).Error(
			fiber.ErrInternalServerError.Code,
			string(orders.InsertOrderErr),
			"user not found",
		).Res()
	}

	// Find Products
	connProduct, err := grpc.Dial(h.cfg.App().ProductsGrpcUrl(), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return entities.NewResponse(c).Error(
			fiber.ErrInternalServerError.Code,
			string(auth.LoginErr),
			err.Error(),
		).Res()
	}
	defer connProduct.Close()

	clientProduct := pbProducts.NewProductsServicesClient(connProduct)

	streamProduct, err := clientProduct.FindOneProduct(ctx)
	if err != nil {
		log.Fatalf("could not stream order with an error: %v", err)
	}
	for _, p := range req.Products {
		if err := streamProduct.Send(&pbProducts.FindOneProdcutReq{Id: p.Product.Id}); err != nil {
			log.Fatalf("%v.Send(%v) = %v", streamProduct, &pbProducts.FindOneProdcutReq{Id: p.Product.Id}, err)
		}
		fmt.Printf("order: %v has been streamed\n", &pbProducts.FindOneProdcutReq{Id: p.Product.Id})
	}
	productsStream, err := streamProduct.CloseAndRecv()
	if err != nil {
		return entities.NewResponse(c).Error(
			fiber.ErrInternalServerError.Code,
			string(orders.InsertOrderErr),
			err.Error(),
		).Res()
	}
	if len(productsStream.Products) == 0 {
		return entities.NewResponse(c).Error(
			fiber.ErrInternalServerError.Code,
			string(orders.InsertOrderErr),
			"prodcuts are empty",
		).Res()
	}

	for i := range productsStream.Products {
		req.Products[i].Product.Id = productsStream.Products[i].Id
		req.Products[i].Product.Title = productsStream.Products[i].Title
		req.Products[i].Product.Description = productsStream.Products[i].Description
		req.Products[i].Product.Price = productsStream.Products[i].Price
		if req.Products[i].Qty == 0 {
			req.Products[i].Qty = 1
		}
		req.Total += productsStream.Products[i].Price * float64(req.Products[i].Qty)
	}

	// Insert Order
	order, err := h.ordersUsecase.InsertOrder(req)
	if err != nil {
		return entities.NewResponse(c).Error(
			fiber.ErrBadRequest.Code,
			string(orders.InsertOrderErr),
			err.Error(),
		).Res()
	}

	return entities.NewResponse(c).Success(
		fiber.StatusOK,
		order,
	).Res()
}
