package order

type Service struct {
	repo Repository
}

func NewOrderService(repo Repository) *Service {
	return &Service{
		repo: repo,
	}
}

func (s *Service) Get(id string) (*Order, error) {
	return s.repo.Get(id)
}
