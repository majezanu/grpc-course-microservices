package main

import (
	"log"

	"github.com/majezanu/grpc-course-microservices/order/config"
	"github.com/majezanu/grpc-course-microservices/payment/internal/adapters/db"
	"github.com/majezanu/grpc-course-microservices/payment/internal/adapters/grpc"
	"github.com/majezanu/grpc-course-microservices/payment/internal/application/core/api"
)

func main() {
	dbAdapter, err := db.NewAdapter(config.GetDataSourceURL())
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	application := api.NewApplication(dbAdapter)
	grpAdapter := grpc.NewAdapter(application, config.GetApplicationPort())
	grpAdapter.Run()
}
