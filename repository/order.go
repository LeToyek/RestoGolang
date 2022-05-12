package repository

import (
	"database/sql"
	"errors"
	"resto/entities"
)

type OrderRepository interface {
	AddOrder(order entities.Order) error
	GetOrder(userID string) (entities.Order, error)
}

func (r *Repository) AddOrder(order entities.Order) error {
	_, err := r.DB.Exec(
		queryInputOrder,
		order.Order_Id,
		order.Order_date,
		order.Created_at,
		order.Updated_at,
		order.User_id,
		order.Table_id,
	)
	return err
}

func (r *Repository) GetOrder(userID string) (entities.Order, error) {
	var order entities.Order

	row := r.DB.QueryRow(queryGetOrder, userID)
	err := row.Scan(
		&order.Order_Id,
		&order.Order_date,
		&order.Created_at,
		&order.Updated_at,
		&order.User_id,
		&order.Table_id,
	)

	if err == sql.ErrNoRows {
		return order, errors.New("no rows were returned")
	}
	if err != nil {
		return order, err
	}

	return order, err
}
