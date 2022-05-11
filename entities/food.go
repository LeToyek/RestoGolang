package entities

type Food struct {
	Food_id    string
	Name       string  `json:"name" validate:"required,min=2,max=100"`
	Price      float64 `json:"price" validate:"required"`
	Image      string  `json:"image" `
	Category   string  `json:"category" validate:"required"`
	Created_at string  `json:"created_at"`
	Updated_at string  `json:"updated_at"`
}
