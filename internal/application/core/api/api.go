package api

import (
	"github.com/aholake/shipping-service/internal/application/core/domain"
	"github.com/aholake/shipping-service/internal/ports"
)

type Application struct {
	dbPort ports.DBPort
}

func NewApplication(dbPort ports.DBPort) *Application {
	return &Application{
		dbPort: dbPort,
	}
}

func (a Application) Ship(shipping domain.Shipping) (domain.Shipping, error) {
	err := a.dbPort.Save(&shipping)
	if err != nil {
		return domain.Shipping{}, err
	}
	return shipping, nil
}
