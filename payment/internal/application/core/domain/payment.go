package domain

import (
	"time"
)

type Payment struct {
	ID         int64   `json:"id"`
	OrderID    int64   `json:"order_id"`
	CustomerID int64   `json:"customer_id"`
	TotalPrice float32 `json:"total_price"`
	CreatedAt  int64   `json:"created_at"`
}

func NewPayment(orderID int64, customerId int64, totalPrice float32) Payment {
	return Payment{
		CreatedAt:  time.Now().Unix(),
		OrderID:    orderID,
		CustomerID: customerId,
		TotalPrice: totalPrice,
	}
}
