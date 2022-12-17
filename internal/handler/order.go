package handler

import (
	"deps_per_service/internal/core"
	"github.com/gofiber/fiber/v2"
)

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

func (h *OrderHandler) GetOrder(ctx *fiber.Ctx) error {
	// TODO implement me
	panic("implement me")
}

func (h *OrderHandler) CreateOrder(ctx *fiber.Ctx) error {
	// TODO implement me
	panic("implement me")
}

func (h *OrderHandler) UpdateOrder(ctx *fiber.Ctx) error {
	// TODO implement me
	panic("implement me")
}

func (h *OrderHandler) DeleteOrder(ctx *fiber.Ctx) error {
	// TODO implement me
	panic("implement me")
}
