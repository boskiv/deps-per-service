package handler

import "deps_per_service/internal/service"

type Handlers struct {
	UserHandler    *UserHandler
	OrderHandler   *OrderHandler
	BillHandler    *BillHandler
	ProfileHandler *ProfileHandler
}

func NewHandlers(
	services *service.Services,
) *Handlers {
	return &Handlers{
		UserHandler:    NewUserHandler(services.UserService),
		OrderHandler:   NewOrderHandler(services.OrderService),
		BillHandler:    NewBillHandler(services.BillService),
		ProfileHandler: NewProfileHandler(services.ProfileService),
	}
}
