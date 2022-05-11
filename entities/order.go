package entities

type Order struct {
	Order_Id   string   `json:"order_id" validate:"required"`
	Order_date string   `json:"order_date" validate:"required"`
	Created_at string   `json:"created_at"`
	Updated_at string   `json:"updated_at"`
	User_id    string   `json:"user_id" validate:"required"`
	Foods_id   []string `json:"foods_id" validate:"required"`
	Table_id   string   `json:"table_id" validate:"required"`
}
