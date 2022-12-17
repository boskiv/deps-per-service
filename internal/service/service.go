package service

import (
	"deps_per_service/internal/repository"
)

type Services struct {
	UserService    *UserService
	OrderService   *OrderService
	BillService    *BillService
	ProfileService *ProfileService
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

func (s *Services) GetServices() *Services {
	return s
}
