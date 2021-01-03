package invoiceItem

import (
	"database/sql"
	"time"
)

//Model of invoiceItem
type Model struct {
	ID              uint
	InvoiceHeaderID uint
	ProductID       uint
	CreatedAt       time.Time
	UpdatedAt       time.Time
}

//Models a group of invoice items
type Models []*Model

// Storage system
type Storage interface {
	Migrate() error
	//tx es transaccion
	CreateTx(*sql.Tx, uint, Models) error
}

//Service of invoiceItem
type Service struct {
	storage Storage
}

//NewService returns a service pointer
func NewService(s Storage) *Service {
	return &Service{s}
}

//Migrate is used to migrate products
func (s *Service) Migrate() error {
	return s.storage.Migrate()
}
