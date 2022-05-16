package repository

import (
	"database/sql"
	"errors"
	"resto/entities"
)

type InvoiceRepository interface {
	AddInvoice(invoice entities.Invoice) error
	GetInvoice(userID string) (entities.Invoice, error)
	GetPrice() ([]int, error)
	GetMoreDetailed(userID string) (entities.Invoice, error)
}

func (r *Repository) AddInvoice(invoice entities.Invoice) error {

	_, err := r.DB.Exec(
		queryAddInvoice,
		invoice.Invoice_id,
		invoice.User_id,
		invoice.Pay_date,
		invoice.Total_cost,
	)

	return err
}

func (r *Repository) GetInvoice(userID string) (entities.Invoice, error) {
	var invoice entities.Invoice

	row := r.DB.QueryRow(queryGetInvoice, userID)

	err := row.Scan(
		&invoice.Invoice_id,
		&invoice.User_id,
		&invoice.Table_number,
		&invoice.Pay_date,
		&invoice.Total_cost,
	)
	if err == sql.ErrNoRows {
		return invoice, errors.New("no rows were returned")
	}
	if err != nil {
		return invoice, err
	}

	return invoice, err
}

func (r *Repository) GetPrice(userID string) ([]int, error) {
	var prices []int

	rows, err := r.DB.Query(queryGetTotalPriceOrder, userID)
	if err != nil && err != sql.ErrNoRows {
		return nil, err
	}

	for rows.Next() {
		var price int
		var err = rows.Scan(
			&price,
		)
		if err != nil {
			return nil, err
		}
		prices = append(prices, price)
	}
	return prices, err
}

func (r *Repository) GetMoreDetailed(userID string) (entities.Invoice, error) {
	var invoice entities.Invoice

	row := r.DB.QueryRow(queryGetMoreDetailedInvoice, userID)

	err := row.Scan(
		&invoice.F_name,
		&invoice.F_count,
		&invoice.Food_price,
		&invoice.Total_price,
	)
	if err == sql.ErrNoRows {
		return invoice, errors.New("no rows were returned")
	}
	if err != nil {
		return invoice, err
	}

	return invoice, err
}
