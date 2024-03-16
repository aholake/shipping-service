package adapters

import (
	"context"

	"github.com/aholake/order-proto/golang/shipping"
	"github.com/aholake/shipping-service/internal/application/core/domain"
	"github.com/aholake/shipping-service/internal/ports"
)

type GRPCAdapter struct {
	apiPort ports.APIPort
	shipping.UnimplementedShippingServer
}

func NewGRPCAdapter(apiPort ports.APIPort) *GRPCAdapter {
	return &GRPCAdapter{
		apiPort: apiPort,
	}
}

func (g GRPCAdapter) Create(ctx context.Context, req *shipping.CreateShippingRequest) (*shipping.CreateShippingResponse, error) {
	_, err := g.apiPort.Ship(domain.Shipping{
		Address: req.Address,
	})
	if err != nil {
		return nil, err
	}
	return &shipping.CreateShippingResponse{}, nil
}
