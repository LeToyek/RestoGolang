package entities

import "time"

type Food struct {
	ID         int
	Name       string    `json:"name" validate:"required,min=2,max=100"`
	Price      float64   `json:"price" validate:"required"`
	Image      string    `json:"image" validate:"required"`
	Category   string    `json:"category" validate:"required"`
	Created_at time.Time `json:"created_at"`
	Updated_at time.Time `json:"updated_at"`
	Food_id    string    `json:"food_id" validate:"required"`
}
