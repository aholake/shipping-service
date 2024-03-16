package ports

import "github.com/aholake/shipping-service/internal/application/core/domain"

type APIPort interface {
	Ship(domain.Shipping) (domain.Shipping, error)
}
