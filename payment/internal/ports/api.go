package ports

import "github.com/majezanu/grpc-course-microservices/payment/internal/application/core/domain"

type APIPort interface {
	Charge(Order domain.Payment) (domain.Payment, error)
}
