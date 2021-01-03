package invoice

import (
	"github.com/CristianPadilla/BDwithGoPostgreSQL/pkg/invoiceHeader"
	"github.com/CristianPadilla/BDwithGoPostgreSQL/pkg/invoiceItem"
)

//Model of invoice
type Model struct {
	Header *invoiceHeader.Model
	Items  invoiceItem.Models
}

// Storage system
type Storage interface {
	//tx es transaccion
	Create(*Model) error
}

//Service of invoice
type Service struct {
	storage Storage
}

//NewService returns a service pointer
func NewService(s Storage) *Service {
	return &Service{s}
}

//Create is used to create products
func (s *Service) Create(m *Model) error {
	return s.storage.Create(m)
}
