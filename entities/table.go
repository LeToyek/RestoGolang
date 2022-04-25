package entities

import "time"

type Table struct {
	ID           int
	Total_guests int       `json:"total_guests" validate:"required"`
	Table_number int       `json:"table_number" validate:"required"`
	Created_at   time.Time `json:"created_at"`
	Updated_at   time.Time `json:"updated_at"`
	Table_id     string    `json:"table_id" validate:"required"`
}
