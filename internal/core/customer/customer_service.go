package customer

import (
	"errors"
	"github.com/google/uuid"
)

type Service struct {
	repository IRepository
}

func NewService(customerRepository IRepository) *Service {
	return &Service{
		repository: customerRepository,
	}
}

func (s *Service) Create(customer Customer) (*Customer, error) {
	existingCustomer, err := s.GetByCpf(customer.Cpf)
	if err != nil {
		return nil, errors.New("failed to get customer by cpf")
	}

	if existingCustomer != nil {
		return nil, errors.New("entered cpf is already registered in our system")
	}

	return s.repository.Create(customer)
}

func (s *Service) Update(customerId uuid.UUID, customer Customer) (*Customer, error) {
	return s.repository.Update(customerId, customer)
}

func (s *Service) Delete(customerId uuid.UUID) error {
	return s.repository.Delete(customerId)
}

func (s *Service) GetAll() ([]Customer, error) {
	return s.repository.GetAll()
}

func (s *Service) GetByCpf(cpf string) (*Customer, error) {
	return s.repository.GetByCpf(cpf)
}
