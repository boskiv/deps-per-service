package handler

import "deps_per_service/internal/core"

type BillService interface {
	GetBill(billID int) (*core.Bill, error)
	CreateBill(bill *core.Bill) error
	UpdateBill(bill *core.Bill) error
	DeleteBill(billID int) error
}

type BillHandler struct {
	billService BillService
}

func NewBillHandler(billService BillService) *BillHandler {
	return &BillHandler{
		billService: billService,
	}
}
