package service

import "resto/entities"

type InvoiceService interface {
	AddInvoice(invoice entities.Invoice) error
	GetInvoice(userID string) (entities.Invoice, error)
	GetPrice(userID string) ([]int, error)
	GetMoreDetailed(userID string) (entities.Invoice, error)
}

func (s *Service) AddInvoice(invoice entities.Invoice) error {
	return s.Repository.AddInvoice(invoice)
}

func (s *Service) GetInvoice(userID string) (entities.Invoice, error) {
	return s.Repository.GetInvoice(userID)
}

func (s *Service) GetPrice(userID string) ([]int, error) {
	return s.Repository.GetPrice(userID)
}
func (s *Service) GetMoreDetailed(userID string) (entities.Invoice, error) {
	return s.Repository.GetMoreDetailed(userID)
}
