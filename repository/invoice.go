package repository

import "resto/entities"

type InvoiceRepository interface {
	AddInvoice(invoice entities.Invoice) error
	GetInvoice(userId string) (entities.Invoice, error)
}

func (r *Repository) AddInvoice(invoice entities.Invoice) error {
	_, err := r.DB.Exec(
		queryAddInvoice,
		invoice.Invoice_id,
		invoice.Order_id,
		invoice.Pay_date,
		invoice.Total_cost,
	)
	return err
}

func (r *Repository) GetInvoice(userId string) {

}
