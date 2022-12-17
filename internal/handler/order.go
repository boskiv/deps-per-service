package handler

import "deps_per_service/internal/core"

type OrderService interface {
	GetOrder(orderID int) (*core.Order, error)
	CreateOrder(order *core.Order) error
	UpdateOrder(order *core.Order) error
	DeleteOrder(orderID int) error
}

type OrderHandler struct {
	orderService OrderService
}

func NewOrderHandler(orderService OrderService) *OrderHandler {
	return &OrderHandler{
		orderService: orderService,
	}
}
