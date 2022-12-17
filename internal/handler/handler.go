package handler

import "deps_per_service/internal/service"

type Services interface {
	GetServices() *service.Services
}

type Handlers struct {
	UserHandler    *UserHandler
	OrderHandler   *OrderHandler
	BillHandler    *BillHandler
	ProfileHandler *ProfileHandler
}

func NewHandlers(
	services Services,
) *Handlers {
	return &Handlers{
		UserHandler:    NewUserHandler(services.GetServices().UserService),
		OrderHandler:   NewOrderHandler(services.GetServices().OrderService),
		BillHandler:    NewBillHandler(services.GetServices().BillService),
		ProfileHandler: NewProfileHandler(services.GetServices().ProfileService),
	}
}
