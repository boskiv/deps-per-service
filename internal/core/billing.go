package core

type Bill struct {
	ID      int
	UserID  int
	OrderID int
	Amount  int
}

func (b *Bill) toServiceDTO() *BillServiceDTO {
	return &BillServiceDTO{
		ID:      b.ID,
		UserID:  b.UserID,
		OrderID: b.OrderID,
		Amount:  b.Amount,
	}
}

// BillServiceDTO is a DTO for BillService. It is used to transfer data between the service and the handler.
type BillServiceDTO struct {
	ID      int
	UserID  int
	OrderID int
	Amount  int
}

func (b *BillServiceDTO) toModel() *Bill {
	return &Bill{
		ID:      b.ID,
		UserID:  b.UserID,
		OrderID: b.OrderID,
		Amount:  b.Amount,
	}
}
