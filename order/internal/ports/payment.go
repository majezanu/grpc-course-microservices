package ports

import "github.com/majezanu/grpc-course-microservices/order/internal/application/core/domain"

type PaymentPort interface {
	Charge(*domain.Order) error
}
