package customer

type Customer struct {
	Id    uint
	Name  string
	Cpf   string
	Email string
	Age   int
}

type IService interface {
	Create(customer Customer) (*Customer, error)
	Update(customerId uint, customer Customer) (*Customer, error)
	Delete(customerId uint) error
	GetAll() ([]Customer, error)
	GetByCpf(customerCpf string) (*Customer, error)
}

type IRepository interface {
	Create(customer Customer) (*Customer, error)
	Update(customerId uint, customer Customer) (*Customer, error)
	Delete(customerId uint) error
	GetAll() ([]Customer, error)
	GetByCpf(customerCpf string) (*Customer, error)
}
