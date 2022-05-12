package service

import "resto/entities"

type OrderService interface {
	AddOrder(order entities.Order) error
	GetOrder(userID string) (entities.Order, error)
}

func (s *Service) AddOrder(order entities.Order) error {
	return s.Repository.AddOrder(order)
}

func (s *Service) GetOrder(userID string) (entities.Order, error) {
	return s.Repository.GetOrder(userID)
}
