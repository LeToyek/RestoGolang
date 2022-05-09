package entities

type Table struct {
	Table_number int    `json:"table_number" validate:"required"`
	Total_guests int    `json:"total_guests" validate:"required"`
	Created_at   string `json:"created_at"`
	Updated_at   string `json:"updated_at"`
}
