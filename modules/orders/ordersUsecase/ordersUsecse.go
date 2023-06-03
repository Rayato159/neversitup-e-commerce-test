package ordersUsecase

import (
	"fmt"

	"github.com/Rayato159/neversuitup-e-commerce-test/modules/orders"
	"github.com/Rayato159/neversuitup-e-commerce-test/modules/orders/ordersRepository"
)

type IOrdersUsecase interface {
	InsertOrder(req *orders.Order) (*orders.Order, error)
	FindOrders(userId string) []*orders.Order
	FindOneOrder(userId, orderId string) (*orders.Order, error)
	CancelOrder(userId, orderId string) (*orders.Order, error)
}

type ordersUsecase struct {
	ordersRepository ordersRepository.IOrdersRepository
}

func NewOrdersUsecase(ordersRepository ordersRepository.IOrdersRepository) IOrdersUsecase {
	return &ordersUsecase{
		ordersRepository: ordersRepository,
	}
}

func (u *ordersUsecase) InsertOrder(req *orders.Order) (*orders.Order, error) {
	orderId, err := u.ordersRepository.InsertOrder(req)
	if err != nil {
		return nil, err
	}

	order, err := u.ordersRepository.FindOneOrder(req.UserId, orderId)
	if err != nil {
		return nil, err
	}
	return order, nil
}

func (u *ordersUsecase) FindOrders(userId string) []*orders.Order {
	return u.ordersRepository.FindOrders(userId)
}

func (u *ordersUsecase) FindOneOrder(userId, orderId string) (*orders.Order, error) {
	order, err := u.ordersRepository.FindOneOrder(userId, orderId)
	if err != nil {
		return nil, err
	}
	return order, nil
}

func (u *ordersUsecase) CancelOrder(userId, orderId string) (*orders.Order, error) {
	status, err := u.ordersRepository.FindOrderStatus(userId, orderId)
	if err != nil {
		return nil, err
	}
	if status != "waiting" {
		return nil, fmt.Errorf("status: %v not available to cancel", err)
	}

	if err := u.ordersRepository.CancelOrder(userId, orderId); err != nil {
		return nil, err
	}

	order, err := u.ordersRepository.FindOneOrder(userId, orderId)
	if err != nil {
		return nil, err
	}
	return order, nil
}
