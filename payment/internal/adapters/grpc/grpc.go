package grpc

import (
	"context"
	"github.com/majezanu/grpc-course-microservices-proto/golang/payment"
	"github.com/majezanu/grpc-course-microservices/payment/internal/application/core/domain"
)

func (adapter Adapter) Create(ctx context.Context, request *payment.CreatePaymentRequest) (*payment.CreatePaymentResponse, error) {
	newPayment := domain.NewPayment(request.OrderId, request.CustomerId, request.TotalPrice)
	result, err := adapter.api.Charge(newPayment)
	if err != nil {
		return nil, err
	}
	return &payment.CreatePaymentResponse{PaymentId: result.ID, OrderId: result.OrderID, CustomerId: result.CustomerID, TotalPrice: result.TotalPrice}, nil
}
