package product

type Service struct {
	repo Repository
}

func NewProductService(repo Repository) *Service {
	return &Service{
		repo: repo,
	}
}

func (s *Service) Search(search string, category string) (*[]Product, error) {
	return s.repo.Search(search, category)
}

func (s *Service) Delete(id string) error {
	return s.repo.Delete(id)
}
