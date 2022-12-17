package service

import (
	"deps_per_service/internal/handler"
	"deps_per_service/internal/repository"
)

type Services struct {
	UserService    handler.UserService
	OrderService   handler.OrderService
	BillService    handler.BillService
	ProfileService handler.ProfileService
}

func NewServices(
	repositories *repository.Repositories,
) *Services {
	return &Services{
		UserService:    NewUserService(repositories.UserRepo),
		OrderService:   NewOrderService(repositories.OrderRepo),
		BillService:    NewBillService(repositories.BillRepo),
		ProfileService: NewProfileService(repositories.ProfileRepo),
	}

}
