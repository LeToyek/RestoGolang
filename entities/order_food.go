package entities

type OrderFoods struct {
	Order_id    string `json:"order_id" validate:"required"`
	Food_id     string `json:"foods_id" validate:"required"`
	Total_count int    `json:"total_count" validate:"required"`
}
