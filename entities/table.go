package entities

type TableResto struct {
	Number     int
	Chairs     int    `json:"total_chairs" validate:"required"`
	IsUsed     bool   `josn:"isUsed" validate:"required"`
	Created_at string `json:"created_at"`
	Updated_at string `json:"updated_at"`
}
