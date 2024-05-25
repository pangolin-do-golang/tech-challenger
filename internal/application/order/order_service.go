package order

import "github.com/google/uuid"

type Service struct {
	repository IRepository
}

func NewOrderService(orderRepository IRepository) *Service {
	return &Service{
		repository: orderRepository,
	}
}

func (s *Service) Get(id uuid.UUID) (*Order, error) {
	return s.repository.Get(id)
}

func (s *Service) GetAll() ([]Order, error) {
	return s.repository.GetAll()
}
