package grpc

import (
	"context"
	"github.com/majezanu/grpc-course-microservices-proto/golang/order"
	"github.com/majezanu/grpc-course-microservices/order/internal/application/core/domain"
)

func (adapter Adapter) Create(ctx context.Context, request *order.CreteOrderRequest) (*order.CreateOrderResponse, error) {
	var orderItems []domain.OrderItem
	for _, orderItem := range request.OrderItems {
		orderItems = append(orderItems, domain.OrderItem{
			ProductCode: orderItem.ProductCode,
			UnitPrice:   orderItem.UnitPrice,
			Quantity:    orderItem.Quantity,
		})
	}
	newOrder := domain.NewOrder(request.CustomerId, orderItems)
	result, err := adapter.api.PlaceOrder(newOrder)
	if err != nil {
		return nil, err
	}
	return &order.CreateOrderResponse{OrderId: result.ID, TotalPrice: result.TotalPrice, TotalQuantity: result.TotalQuantity}, nil
}
