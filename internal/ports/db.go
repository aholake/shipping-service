package ports

import "github.com/aholake/shipping-service/internal/application/core/domain"

type DBPort interface {
	Save(*domain.Shipping) error
}
