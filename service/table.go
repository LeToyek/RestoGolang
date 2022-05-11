package service

import "resto/entities"

type TableRestoService interface {
	AddTable(table entities.TableResto) error
	GetAllTables() ([]entities.TableResto, error)
}

func (s *Service) AddTable(table entities.TableResto) error {
	err := s.Repository.AddTable(table)
	return err
}
func (s *Service) GetAllTables() ([]entities.TableResto, error) {
	tables, err := s.Repository.GetAllTables()
	return tables, err
}
