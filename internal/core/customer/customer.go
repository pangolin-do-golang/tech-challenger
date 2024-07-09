package customer

import "github.com/google/uuid"

type Customer struct {
	Id    uuid.UUID `json:"id"`
	Name  string    `json:"name"`
	Cpf   string    `json:"cpf"`
	Email string    `json:"email"`
	Age   int       `json:"age"`
}

type IService interface {
	Create(customer Customer) (*Customer, error)
	Update(customerId uuid.UUID, customer Customer) (*Customer, error)
	Delete(customerId uuid.UUID) error
	GetAll() ([]Customer, error)
	GetByCpf(customerCpf string) (*Customer, error)
}

type IRepository interface {
	Create(customer Customer) (*Customer, error)
	Update(customerId uuid.UUID, customer Customer) (*Customer, error)
	Delete(customerId uuid.UUID) error
	GetAll() ([]Customer, error)
	GetByCpf(customerCpf string) (*Customer, error)
}
