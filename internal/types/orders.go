package types

import "time"

type (
	Order struct {
		OrderID   string    `json:"order_id"`
		Status    string    `json:"status"`
		Total     Money     `json:"total"`
		CreatedAt time.Time `json:"created_at"`
	}

	OrderStatus string
)
