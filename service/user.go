package service

import "resto/entities"

type UserService interface {
	AddUser(user entities.User) (entities.User, error)
	// GetUsers(users []entities.User) ([]entities.User, error)
	// GetUser(user entities.User) (entities.User, error)
}

func (s *Service) AddUser(user entities.User) (entities.User, error) {
	return s.Repository.AddUser(user)
}
