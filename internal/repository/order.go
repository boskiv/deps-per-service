package repository

import (
	"deps_per_service/internal/core"
	"gorm.io/gorm"
)

type OrderRepo struct {
	BaseRepo
}

func (o *OrderRepo) GetOrder(orderID int) (*core.Order, error) {
	//TODO implement me
	panic("implement me")
}

func (o *OrderRepo) CreateOrder(order *core.Order) error {
	//TODO implement me
	panic("implement me")
}

func (o *OrderRepo) UpdateOrder(order *core.Order) error {
	//TODO implement me
	panic("implement me")
}

func (o *OrderRepo) DeleteOrder(orderID int) error {
	//TODO implement me
	panic("implement me")
}

func NewOrderRepo(db *gorm.DB) *OrderRepo {
	return &OrderRepo{
		BaseRepo: BaseRepo{
			db: db,
		},
	}
}
