package main

import (
	"log"

	"github.com/majezanu/grpc-course-microservices/order/config"
	"github.com/majezanu/grpc-course-microservices/order/internal/adapters/db"
	"github.com/majezanu/grpc-course-microservices/order/internal/adapters/grpc"
	"github.com/majezanu/grpc-course-microservices/order/internal/adapters/payment"
	"github.com/majezanu/grpc-course-microservices/order/internal/application/core/api"
)

func main() {
	dbAdapter, err := db.NewAdapter(config.GetDataSourceURL())
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	paymentAdapter, err := payment.NewAdapter(config.GetPaymentServiceURL())
	if err != nil {
		log.Fatalf("Failed to connect to payment service: %v", err)
	}
	application := api.NewApplication(dbAdapter, paymentAdapter)
	grpAdapter := grpc.NewAdapter(application, config.GetApplicationPort())
	grpAdapter.Run()
}
