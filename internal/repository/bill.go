package repository

import (
	"deps_per_service/internal/core"
	"gorm.io/gorm"
)

type BillRepo struct {
	BaseRepo
}

func (b *BillRepo) GetBill(billID int) (*core.BillServiceDTO, error) {
	//TODO implement me
	panic("implement me")
}

func (b *BillRepo) CreateBill(bill *core.BillServiceDTO) error {
	//TODO implement me
	panic("implement me")
}

func (b *BillRepo) UpdateBill(bill *core.BillServiceDTO) error {
	//TODO implement me
	panic("implement me")
}

func (b *BillRepo) DeleteBill(billID int) error {
	//TODO implement me
	panic("implement me")
}

func NewBillRepo(db *gorm.DB) *BillRepo {
	return &BillRepo{
		BaseRepo: BaseRepo{
			db: db,
		},
	}
}
