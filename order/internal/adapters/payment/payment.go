package payment

import (
	"context"

	"github.com/majezanu/grpc-course-microservices-proto/golang/payment"
	"github.com/majezanu/grpc-course-microservices/order/internal/application/core/domain"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type Adapter struct {
	payment payment.PaymentClient
}

func NewAdapter(paymentServiceURL string) (*Adapter, error) {
	var opts []grpc.DialOption
	opts = append(opts, grpc.WithTransportCredentials(insecure.NewCredentials()))
	conn, err := grpc.NewClient(paymentServiceURL, opts...)
	if err != nil {
		return nil, err
	}

	client := payment.NewPaymentClient(conn)
	return &Adapter{
		payment: client,
	}, nil
}

func (adapter *Adapter) Charge(order *domain.Order) error {
	_, err := adapter.payment.Create(context.Background(),
		&payment.CreatePaymentRequest{
			CustomerId: order.CustomerID,
			OrderId:    order.ID,
			TotalPrice: order.TotalPrice,
		})
	return err
}
