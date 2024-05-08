package order

import "github.com/google/uuid"

type Service struct {
	repo Repository
}

func NewOrderService(repo Repository) *Service {
	return &Service{
		repo: repo,
	}
}

func (s *Service) Get(id uuid.UUID) (*Order, error) {
	return s.repo.Get(id)
}
