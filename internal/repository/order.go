package repository

import (
	"deps_per_service/internal/core"
	"deps_per_service/internal/service"
	"gorm.io/gorm"
)

type orderRepo struct {
	BaseRepo
}

func (o *orderRepo) GetOrder(orderID int) (*core.Order, error) {
	//TODO implement me
	panic("implement me")
}

func (o *orderRepo) CreateOrder(order *core.Order) error {
	//TODO implement me
	panic("implement me")
}

func (o *orderRepo) UpdateOrder(order *core.Order) error {
	//TODO implement me
	panic("implement me")
}

func (o *orderRepo) DeleteOrder(orderID int) error {
	//TODO implement me
	panic("implement me")
}

func NewOrderRepo(db *gorm.DB) service.OrderRepo {
	return &orderRepo{
		BaseRepo: BaseRepo{
			db: db,
		},
	}
}
