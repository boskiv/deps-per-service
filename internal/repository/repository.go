package repository

import (
	"gorm.io/gorm"
)

type BaseRepo struct {
	db *gorm.DB
}

type Repositories struct {
	UserRepo    *UserRepo
	ProfileRepo *ProfileRepo
	OrderRepo   *OrderRepo
	BillRepo    *BillRepo
}

func NewRepositories(db *gorm.DB) *Repositories {
	return &Repositories{
		UserRepo:    NewUserRepo(db),
		ProfileRepo: NewProfileRepo(db),
		OrderRepo:   NewOrderRepo(db),
		BillRepo:    NewBillRepo(db),
	}
}
