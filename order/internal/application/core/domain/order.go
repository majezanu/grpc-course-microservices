package domain

import (
	"time"
)

type OrderItem struct {
	ProductCode string  `json:"product_code"`
	UnitPrice   float32 `json:"unit_price"`
	Quantity    int32   `json:"quantity"`
}

type Order struct {
	ID            int64       `json:"id"`
	CustomerID    int64       `json:"customer_id"`
	Status        string      `json:"status"`
	TotalPrice    float32     `json:"total_price"`
	TotalQuantity int32       `json:"total_quantity"`
	OrderItems    []OrderItem `json:"order_items"`
	CreatedAt     int64       `json:"created_at"`
}

func NewOrder(customerId int64, orderItems []OrderItem) Order {
	var totalPrice float32
	var totalQuantity int32
	for _, item := range orderItems {
		totalPrice += item.UnitPrice * float32(item.Quantity)
		totalQuantity += item.Quantity
	}
	return Order{
		CreatedAt:     time.Now().Unix(),
		Status:        "Pending",
		CustomerID:    customerId,
		OrderItems:    orderItems,
		TotalPrice:    totalPrice,
		TotalQuantity: totalQuantity,
	}
}
