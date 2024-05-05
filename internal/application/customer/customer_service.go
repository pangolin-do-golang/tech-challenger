package customer

type Service struct {
	repository IRepository
}

func NewService(customerRepository IRepository) *Service {
	return &Service{
		repository: customerRepository,
	}
}

func (s *Service) Create(customer Customer) (*Customer, error) {
	return s.repository.Create(customer)
}

func (s *Service) Update(customerId uint, customer Customer) (*Customer, error) {
	return s.repository.Update(customerId, customer)
}

func (s *Service) Delete(customerId uint) error {
	return s.repository.Delete(customerId)
}

func (s *Service) GetAll() ([]Customer, error) {
	return s.repository.GetAll()
}

func (s *Service) GetByCpf(cpf string) (*Customer, error) {
	return s.repository.GetByCpf(cpf)
}
