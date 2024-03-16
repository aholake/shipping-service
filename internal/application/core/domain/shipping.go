package domain

import "time"

type Shipping struct {
	ID        int64  `json:"id"`
	OrderID   int64  `json:"order_id"`
	Address   string `json:"address"`
	CreatedAt int64  `json:"created_at"`
}

func NewShipping(orderId int64, address string) Shipping {
	return Shipping{
		OrderID:   orderId,
		Address:   address,
		CreatedAt: time.Now().Unix(),
	}
}
