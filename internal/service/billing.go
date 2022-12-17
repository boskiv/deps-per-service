package service

import (
	"deps_per_service/internal/core"
)

type BillRepo interface {
	GetBill(billID int) (*core.BillServiceDTO, error)
	CreateBill(bill *core.BillServiceDTO) error
	UpdateBill(bill *core.BillServiceDTO) error
	DeleteBill(billID int) error
}

type BillService struct {
	billRepo BillRepo
}

func NewBillService(billRepo BillRepo) *BillService {
	return &BillService{
		billRepo: billRepo,
	}
}

func (b *BillService) GetBill(billID int) (*core.Bill, error) {
	//TODO implement me
	panic("implement me")
}

func (b *BillService) CreateBill(bill *core.Bill) error {
	//TODO implement me
	panic("implement me")
}

func (b *BillService) UpdateBill(bill *core.Bill) error {
	//TODO implement me
	panic("implement me")
}

func (b *BillService) DeleteBill(billID int) error {
	//TODO implement me
	panic("implement me")
}
