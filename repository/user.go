package repository

import (
	"database/sql"
	"errors"
	"resto/entities"
)

type UserRepository interface {
	AddUser(user entities.User) (entities.User, error)
	GetUsers(users []entities.User) ([]entities.User, error)
	// GetUser(user entities.User) (entities.User, error)
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
