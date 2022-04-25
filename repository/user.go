package repository

import (
	"resto/entities"
)

type UserRepository interface {
	AddUser(user entities.User) (entities.User, error)
	// GetUsers(users []entities.User) ([]entities.User, error)
	// GetUser(user entities.User) (entities.User, error)
}

func (r *Repository) AddUser(user entities.User) (entities.User, error) {
	_, err := r.DB.Exec(
		queryInsertUser,
		user.First_name,
		user.Last_name,
		user.Email,
		user.Password,
		user.Phone,
		user.Token,
		user.Refresh_Token,
		user.User_id,
	)
	return user, err
}
