package invoiceHeader

import (
	"database/sql"
	"time"
)

//Model of invoiceHeader
type Model struct {
	ID        uint
	Client    string
	CreatedAt time.Time
	UpdatedAt time.Time
}

// Storage system
type Storage interface {
	Migrate() error
	//tx es transaccion
	CreateTx(*sql.Tx, *Model) error
}

//Service of invoiceHeader
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
