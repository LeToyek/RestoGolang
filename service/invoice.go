package service

import "resto/entities"

type InvoiceService interface {
	AddInvoice(invoice entities.Invoice) error
	GetInvoice(OrderID string) (entities.Invoice, error)
	GetPrice(OrderID string) ([]int, error)
	GetMoreDetailed(OrderID string) (entities.Invoice, error)
	GetOrderID(userID string) ([]string, error)
}

func (s *Service) AddInvoice(invoice entities.Invoice) error {
	return s.Repository.AddInvoice(invoice)
}

func (s *Service) GetInvoice(OrderID string) (entities.Invoice, error) {
	return s.Repository.GetInvoice(OrderID)
}

func (s *Service) GetPrice(OrderID string) ([]int, error) {
	return s.Repository.GetPrice(OrderID)
}
func (s *Service) GetMoreDetailed(OrderID string) (entities.Invoice, error) {
	return s.Repository.GetMoreDetailed(OrderID)
}
func (s *Service) GetOrderID(userID string) ([]string, error) {
	return s.Repository.GetOrderID(userID)
}
