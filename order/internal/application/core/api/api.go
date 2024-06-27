package api

import (
	"github.com/majezanu/grpc-course-microservices/order/internal/application/core/domain"
	"github.com/majezanu/grpc-course-microservices/order/internal/ports"
)

type Application struct {
	db ports.DBPort
}

func NewApplication(db ports.DBPort) *Application {
	return &Application{
		db: db,
	}
}

func (a Application) PlaceOrder(order domain.Order) (domain.Order, error) {
	for _, item := range order.OrderItems {
		order.TotalPrice += item.UnitPrice * float32(item.Quantity)
		order.TotalQuantity += item.Quantity
	}

	err := a.db.Save(&order)
	if err != nil {
		return domain.Order{}, err
	}
	return order, nil
}
