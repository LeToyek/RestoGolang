package service

import "resto/entities"

type FoodsService interface {
	AddFood(food entities.Food) error
	GetAllfoods() ([]entities.Food, error)
	GetFoodsByCategory(category string) ([]entities.Food, error)
}

func (s *Service) AddFood(food entities.Food) error {
	return s.Repository.AddFood(food)
}

func (s *Service) GetAllfoods() ([]entities.Food, error) {
	return s.Repository.GetAllfoods()
}

func (s *Service) GetFoodsByCategory(category string) ([]entities.Food, error) {
	return s.Repository.GetFoodsByCategory(category)
}
