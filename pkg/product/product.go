package product

import (
	"errors"
	"fmt"
	"time"
)

var errIDNotFound = errors.New("el producto no tiene un id")

//Model of product
type Model struct {
	ID           uint
	Name         string
	Observations string
	Price        int
	CreatedAt    time.Time
	UpdatedAt    time.Time
}	

func (m *Model) String() string {
	return fmt.Sprintf("%02d | %-20s | %-20s | %5d | %10s | %10s ", m.ID, m.Name, m.Observations,
		m.Price, m.CreatedAt.Format("2006-01-02"), m.UpdatedAt.Format("2006-01-02"))
}

//Models is a slice of model
type Models []*Model

// Storage if a interface that have crud methods
type Storage interface {
	Migrate() error
	Create(*Model) error
	Update(*Model) error
	GetAll() (Models, error)
	GetByID(uint) (*Model, error)
	Delete(uint) error
}

//Service of product
type Service struct {
	//el servicio tiene un sistema de almacenamiento
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

//Create is used to create products
func (s *Service) Create(m *Model) error {
	//asignar fecha de creacion del producto
	m.CreatedAt = time.Now()
	return s.storage.Create(m)
}

//GetAll is used to get all products
func (s *Service) GetAll() (Models, error) {
	//asignar fecha de creacion del producto
	return s.storage.GetAll()
}

//GetByID is used to get one product by its id
func (s *Service) GetByID(id uint) (*Model, error) {
	return s.storage.GetByID(id)
}

//Update is used to update one product by its id
func (s *Service) Update(m *Model) error {
	if m.ID == 0 {
		return errIDNotFound
	}
	m.UpdatedAt = time.Now()
	return s.storage.Update(m)
}

//Delete is used to delete one product by its id
func (s *Service) Delete(id uint) error {
	return s.storage.Delete(id)
}
