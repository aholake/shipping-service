package adapters

import (
	"fmt"

	"github.com/aholake/shipping-service/internal/application/core/domain"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type DBAdapter struct {
	db *gorm.DB
}

func NewAdapter(connectionUrl string) (*DBAdapter, error) {
	db, err := gorm.Open(mysql.Open(connectionUrl), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	err = db.AutoMigrate(&Shipping{})
	if err != nil {
		return nil, fmt.Errorf("unable to migrate shipping table, error: %v", err)
	}
	return &DBAdapter{
		db: db,
	}, nil
}

type Shipping struct {
	gorm.Model
	OrderID int64
	Address string
}

func (a DBAdapter) Save(shipping *domain.Shipping) error {
	shippingEntity := Shipping{
		OrderID: shipping.OrderID,
		Address: shipping.Address,
	}
	tx := a.db.Create(&shippingEntity)
	if tx.Error == nil {
		shipping.ID = int64(shippingEntity.ID)
	}
	return tx.Error
}
