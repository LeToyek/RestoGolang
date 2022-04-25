package entities

type Invoice struct {
	ID         int
	Invoice_id string `json:"invoice_id"`
	Order_id   string `json:"order_id"`
}
