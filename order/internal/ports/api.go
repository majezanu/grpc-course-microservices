package ports

import "github.com/majezanu/grpc-course-microservices/order/internal/application/core/domain"

type APIPort interface {
	PlaceOrder(Order domain.Order) (domain.Order, error)
}
