package ports

import "github.com/majezanu/grpc-course-microservices/payment/internal/application/core/domain"

type DBPort interface {
	Get(id string) (domain.Payment, error)
	Save(*domain.Payment) error
}
