package repository

import (
	"deps_per_service/internal/core"
	"deps_per_service/internal/service"
	"gorm.io/gorm"
)

type billRepo struct {
	BaseRepo
}

func (b *billRepo) GetBill(billID int) (*core.BillServiceDTO, error) {
	//TODO implement me
	panic("implement me")
}

func (b *billRepo) CreateBill(bill *core.BillServiceDTO) error {
	//TODO implement me
	panic("implement me")
}

func (b *billRepo) UpdateBill(bill *core.BillServiceDTO) error {
	//TODO implement me
	panic("implement me")
}

func (b *billRepo) DeleteBill(billID int) error {
	//TODO implement me
	panic("implement me")
}

func NewBillRepo(db *gorm.DB) service.BillRepo {
	return &billRepo{
		BaseRepo: BaseRepo{
			db: db,
		},
	}
}
