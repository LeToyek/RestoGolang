package entities

type Invoice struct {
	Invoice_id   string         `json:"invoice_id"`
	Order_ID     string         `json:"order_id"`
	Table_number int            `json:"table_number"`
	Pay_date     string         `json:"Date"`
	Food_detail  []DummyInvoice `json:"Food_detail"`
	Total_cost   float64        `json:"Total_cost"`
}
type DummyInvoice struct {
	F_name       string `json:"name"`
	F_count      int    `json:"count"`
	F_price      int    `json:"price"`
	F_totalPrice int    `json:"total price"`
}
