package ordersUsecase

import (
	"github.com/Rayato159/neversuitup-e-commerce-test/modules/orders"
	"github.com/Rayato159/neversuitup-e-commerce-test/modules/orders/ordersRepository"
)

type IOrdersUsecase interface {
	InsertOrder(req *orders.Order) error
	FindOrders(userId string) []*orders.Order
	FindOneOrder(userId string, orderId string) (*orders.Order, error)
	CancelOrder(userId string) (*orders.Order, error)
}

type ordersUsecase struct {
	ordersRepository ordersRepository.IOrdersRepository
}

func NewOrdersUsecase(ordersRepository ordersRepository.IOrdersRepository) IOrdersUsecase {
	return &ordersUsecase{
		ordersRepository: ordersRepository,
	}
}

func (u *ordersUsecase) InsertOrder(req *orders.Order) error { return nil }

func (u *ordersUsecase) FindOrders(userId string) []*orders.Order {
	return u.ordersRepository.FindOrders(userId)
}

func (u *ordersUsecase) FindOneOrder(userId string, orderId string) (*orders.Order, error) {
	order, err := u.ordersRepository.FindOneOrder(userId, orderId)
	if err != nil {
		return nil, err
	}
	return order, nil
}

func (u *ordersUsecase) CancelOrder(userId string) (*orders.Order, error) { return nil, nil }
