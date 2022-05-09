package repository

import (
	"database/sql"
	"errors"
	"resto/entities"
)

type UserRepository interface {
	AddUser(user entities.User) error
	GetUsers() ([]entities.User, error)
	GetUser(id string) (entities.User, error)
	GetLogin(email string, password string) error
}

func (r *Repository) AddUser(user entities.User) error {
	_, err := r.DB.Exec(
		queryInsertUser,
		user.First_name,
		user.Last_name,
		user.Email,
		user.Password,
		user.Phone,
		user.Created_at,
		user.Updated_at,
		user.User_id,
	)
	return err
}

func (r *Repository) GetUsers() ([]entities.User, error) {
	var users []entities.User

	rows, err := r.DB.Query(queryGetAll)
	if err != nil && err != sql.ErrNoRows {
		return nil, err
	}

	for rows.Next() {
		var user = entities.User{}
		var err = rows.Scan(
			&user.First_name,
			&user.Last_name,
			&user.Email,
			&user.Password,
			&user.Phone,
			&user.Created_at,
			&user.Updated_at,
			&user.User_id,
		)
		if err != nil {
			panic(err)
		}

		users = append(users, user)
	}
	return users, err
}
func (r *Repository) GetUser(id string) (entities.User, error) {
	var user entities.User

	row := r.DB.QueryRow(queryGetOne, id)
	err := row.Scan(
		&user.First_name,
		&user.Last_name,
		&user.Email,
		&user.Password,
		&user.Phone,
		&user.Created_at,
		&user.Updated_at,
		&user.User_id,
	)
	if err == sql.ErrNoRows {
		return user, errors.New("no rows were returned")
	}
	if err != nil {
		return user, err
	}

	return user, err
}
func (r *Repository) GetLogin(email string, password string) error {
	var user entities.User

	row := r.DB.QueryRow(queryGetAcc, email)
	err := row.Scan(
		&user.Password,
	)
	if err == sql.ErrNoRows {
		return errors.New("email tidak ditemukan")
	}
	if password != user.Password {
		return errors.New("wrong password")
	}
	if err != nil {
		return err
	}
	return err
}
