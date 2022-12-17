package service

import (
	"deps_per_service/internal/core"
)

type OrderRepo interface {
	GetOrder(orderID int) (*core.Order, error)
	CreateOrder(order *core.Order) error
	UpdateOrder(order *core.Order) error
	DeleteOrder(orderID int) error
}

type OrderService struct {
	orderRepo OrderRepo
}

func (o *OrderService) GetOrder(orderID int) (*core.Order, error) {
	//TODO implement me
	panic("implement me")
}

func (o *OrderService) CreateOrder(order *core.Order) error {
	//TODO implement me
	panic("implement me")
}

func (o *OrderService) UpdateOrder(order *core.Order) error {
	//TODO implement me
	panic("implement me")
}

func (o *OrderService) DeleteOrder(orderID int) error {
	//TODO implement me
	panic("implement me")
}

func NewOrderService(orderRepo OrderRepo) *OrderService {
	return &OrderService{
		orderRepo: orderRepo,
	}
}
