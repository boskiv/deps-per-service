package service

import (
	"deps_per_service/internal/core"
	"deps_per_service/internal/handler"
)

type BillRepo interface {
	GetBill(billID int) (*core.BillServiceDTO, error)
	CreateBill(bill *core.BillServiceDTO) error
	UpdateBill(bill *core.BillServiceDTO) error
	DeleteBill(billID int) error
}

type billService struct {
	billRepo BillRepo
}

func (b *billService) GetBill(billID int) (*core.Bill, error) {
	//TODO implement me
	panic("implement me")
}

func (b *billService) CreateBill(bill *core.Bill) error {
	//TODO implement me
	panic("implement me")
}

func (b *billService) UpdateBill(bill *core.Bill) error {
	//TODO implement me
	panic("implement me")
}

func (b *billService) DeleteBill(billID int) error {
	//TODO implement me
	panic("implement me")
}

func NewBillService(billRepo BillRepo) handler.BillService {
	return &billService{
		billRepo: billRepo,
	}
}
