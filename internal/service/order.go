package service

import (
	"deps_per_service/internal/core"
	"deps_per_service/internal/handler"
)

type OrderRepo interface {
	GetOrder(orderID int) (*core.Order, error)
	CreateOrder(order *core.Order) error
	UpdateOrder(order *core.Order) error
	DeleteOrder(orderID int) error
}

type orderService struct {
	orderRepo OrderRepo
}

func (o *orderService) GetOrder(orderID int) (*core.Order, error) {
	//TODO implement me
	panic("implement me")
}

func (o *orderService) CreateOrder(order *core.Order) error {
	//TODO implement me
	panic("implement me")
}

func (o *orderService) UpdateOrder(order *core.Order) error {
	//TODO implement me
	panic("implement me")
}

func (o *orderService) DeleteOrder(orderID int) error {
	//TODO implement me
	panic("implement me")
}

func NewOrderService(orderRepo OrderRepo) handler.OrderService {
	return &orderService{
		orderRepo: orderRepo,
	}
}
