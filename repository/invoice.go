package repository

import (
	"database/sql"
	"errors"
	"resto/entities"
)

type InvoiceRepository interface {
	AddInvoice(invoice entities.Invoice) error
	GetInvoice(OrderID string) (entities.Invoice, error)
	GetPrice() ([]int, error)
	GetMoreDetailed(OrderID string) ([]entities.DummyInvoice, error)
	GetOrderID(userID string) ([]string, error)
}

func (r *Repository) AddInvoice(invoice entities.Invoice) error {

	_, err := r.DB.Exec(
		queryAddInvoice,
		invoice.Invoice_id,
		invoice.Order_ID,
		invoice.Pay_date,
		invoice.Total_cost,
	)

	return err
}

func (r *Repository) GetInvoice(OrderID string) (entities.Invoice, error) {
	var invoice entities.Invoice

	row := r.DB.QueryRow(queryGetInvoice, OrderID)

	err := row.Scan(
		&invoice.Invoice_id,
		&invoice.Order_ID,
		&invoice.Pay_date,
		&invoice.Total_cost,
		&invoice.Table_number,
	)
	if err == sql.ErrNoRows {
		return invoice, errors.New("no rows were returned")
	}
	if err != nil {
		return invoice, err
	}

	return invoice, err
}

func (r *Repository) GetOrderID(userID string) ([]string, error) {
	var orders []string

	rows, err := r.DB.Query(queryGetUserOrderID, userID)
	if err != nil && err != sql.ErrNoRows {
		return nil, err
	}

	for rows.Next() {
		var order string
		var err = rows.Scan(
			&order,
		)
		if err != nil {
			return nil, err
		}
		orders = append(orders, order)
	}
	return orders, err
}

func (r *Repository) GetPrice(OrderID string) ([]int, error) {
	var prices []int

	rows, err := r.DB.Query(queryGetTotalPriceOrder, OrderID)
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

func (r *Repository) GetMoreDetailed(OrderID string) ([]entities.DummyInvoice, error) {
	var details []entities.DummyInvoice

	rows, err := r.DB.Query(queryGetMoreDetailedInvoice, OrderID)
	if err != nil && err != sql.ErrNoRows {
		return nil, err
	}

	for rows.Next() {
		var detail entities.DummyInvoice
		var err = rows.Scan(
			&detail.F_name,
			&detail.F_count,
			&detail.F_price,
			&detail.F_totalPrice,
		)
		if err != nil {
			return nil, err
		}
		details = append(details, detail)
	}

	return details, err
}
