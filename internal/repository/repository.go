package repository

import (
	"deps_per_service/internal/service"
	"gorm.io/gorm"
)

type BaseRepo struct {
	db *gorm.DB
}

type Repositories struct {
	UserRepo    service.UserRepo
	ProfileRepo service.ProfileRepo
	OrderRepo   service.OrderRepo
	BillRepo    service.BillRepo
}

func NewRepositories(db *gorm.DB) *Repositories {
	return &Repositories{
		UserRepo:    NewUserRepo(db),
		ProfileRepo: NewProfileRepo(db),
		OrderRepo:   NewOrderRepo(db),
		BillRepo:    NewBillRepo(db),
	}
}
