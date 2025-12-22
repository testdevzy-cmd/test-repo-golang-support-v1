package models

import "time"

type Order struct {
	ID          int
	UserID      int
	ProductIDs  []int
	TotalAmount float64
	CreatedAt   time.Time
}

func NewOrder(id, userID int, productIDs []int, total float64) *Order {
	return &Order{
		ID:          id,
		UserID:      userID,
		ProductIDs:  productIDs,
		TotalAmount: total,
		CreatedAt:   time.Now(),
	}
}

func (o *Order) ProductCount() int {
	return len(o.ProductIDs)
}
