package adapters

import (
	"fmt"
	"log"
	"net"

	"github.com/aholake/order-proto/golang/shipping"
	"github.com/aholake/shipping-service/internal/ports"
	"google.golang.org/grpc"
)

type Server struct {
	port    int32
	apiPort ports.APIPort
}

func NewServer(port int32, apiPort ports.APIPort) *Server {
	return &Server{
		port:    port,
		apiPort: apiPort,
	}
}

func (s Server) Run() {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", s.port))
	if err != nil {
		log.Fatalf("unable to listen on port %d", s.port)
	}

	grpcServer := grpc.NewServer()
	shipping.RegisterShippingServer(grpcServer, NewGRPCAdapter(s.apiPort))
	if err = grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to serve grpc on port %d, error: %v", s.port, err)
	}
}
