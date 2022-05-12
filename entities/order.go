package entities

type Order struct {
	Order_Id   string `json:"order_id"`
	Order_date string `json:"order_date" validate:"required"`
	Created_at string `json:"created_at"`
	Updated_at string `json:"updated_at"`
	User_id    string `json:"user_id"`
	Table_id   int    `json:"table_id" validate:"required"`
}
