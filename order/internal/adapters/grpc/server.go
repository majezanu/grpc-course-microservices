package grpc

import (
	"fmt"
	"log"
	"net"

	"github.com/majezanu/grpc-course-microservices-proto/golang/order"
	"github.com/majezanu/grpc-course-microservices/order/config"
	"github.com/majezanu/grpc-course-microservices/order/internal/ports"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type Adapter struct {
	api  ports.APIPort
	port int
	order.UnimplementedOrderServer
}

func NewAdapter(api ports.APIPort, port int) *Adapter {
	return &Adapter{
		api:  api,
		port: port,
	}
}

func (adapter Adapter) Run() {
	var err error

	listen, err := net.Listen("tcp", fmt.Sprintf(":%d", adapter.port))
	if err != nil {
		log.Fatalf("Failed to list on port %d, error: %v", adapter.port, err)
	}
	grpcServer := grpc.NewServer()
	order.RegisterOrderServer(grpcServer, adapter)
	if config.GetEnv() == "development" {
		reflection.Register(grpcServer)
	}
	if err := grpcServer.Serve(listen); err != nil {
		log.Fatalf("Failed to serve grpc on port")
	}
}
