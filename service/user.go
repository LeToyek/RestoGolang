package service

import "resto/entities"

type UserService interface {
	AddUser(user entities.User) error
	GetUsers() ([]entities.User, error)
	GetUser(id string) (entities.User, error)
	GetLogin(email string, password string) (string, error)
}

func (s *Service) AddUser(user entities.User) error {
	return s.Repository.AddUser(user)
}
func (s *Service) GetUsers() ([]entities.User, error) {
	return s.Repository.GetUsers()
}
func (s *Service) GetUser(id string) (entities.User, error) {
	return s.Repository.GetUser(id)
}
func (s *Service) GetLogin(email string, password string) (string, error) {
	return s.Repository.GetLogin(email, password)
}
