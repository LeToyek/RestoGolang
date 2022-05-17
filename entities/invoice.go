package entities

type Invoice struct {
	Invoice_id   string   `json:"invoice_id"`
	Order_ID     string   `json:"order_id"`
	Table_number int      `json:"table_number"`
	Pay_date     string   `json:"Date"`
	Food_name    []string `json:"Food_name"`
	Food_count   []int    `json:"Food_count"`
	Food_price   int      `json:"Food_price"`
	Total_price  int      `json:"Total_price"`
	Total_cost   float64  `json:"Total_cost"`
	F_name       string
	F_count      int
}
