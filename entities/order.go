package entities

import "time"

type Order struct {
	ID         int
	Order_date time.Time `json:"order_date" validate:"required"`
	Created_at time.Time `json:"created_at"`
	Updated_at time.Time `json:"updated_at"`
	Order_id   string    `json:"order_id" validate:"required"`
	Foods_id   []string  `json:"foods_id" validate:"required"`
	Table_id   string    `json:"table_id" validate:"required"`
}
