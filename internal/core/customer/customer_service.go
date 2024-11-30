package customer

import (
	"errors"
	"github.com/go-resty/resty/v2"
	"github.com/google/uuid"
	"os"
)

type Service struct {
	HttpClient resty.Client
}

func NewService() *Service {
	client := resty.New()
	client.SetBaseURL(os.Getenv("CUSTOMER_SERVICE_URL"))
	return &Service{
		HttpClient: *client,
	}
}

type CreateCustomerPayload struct {
	Name  string `json:"name" binding:"required,min=5,max=20"`
	Cpf   string `json:"cpf" binding:"required,numeric,len=11"`
	Email string `json:"email" binding:"required,email"`
	Age   int    `json:"age" binding:"gte=18,lte=120"`
}

func (s *Service) Create(customer Customer) (*Customer, error) {
	existingCustomer, _ := s.GetByCpf(customer.Cpf)

	if existingCustomer != nil {
		return nil, errors.New("entered cpf is already registered in our system")
	}

	var c *Customer

	_, err := s.HttpClient.R().
		SetBody(CreateCustomerPayload{
			Name:  customer.Name,
			Cpf:   customer.Cpf,
			Email: customer.Email,
			Age:   customer.Age,
		}).
		SetResult(&c).
		Post("/customer")

	if err != nil {
		return nil, err
	}

	return c, nil
}

func (s *Service) Update(customerId uuid.UUID, customer Customer) (*Customer, error) {
	var c *Customer

	_, err := s.HttpClient.R().
		SetPathParams(map[string]string{
			"id": customerId.String(),
		}).
		SetResult(&c).
		Put("/customer/{id}")

	if err != nil {
		return nil, err
	}

	return c, nil
}

func (s *Service) Delete(customerId uuid.UUID) error {
	_, err := s.HttpClient.R().
		SetPathParams(map[string]string{
			"id": customerId.String(),
		}).
		Get("/customer/{id}")

	if err != nil {
		return err
	}

	return nil
}

func (s *Service) GetAll() ([]Customer, error) {
	var customers []Customer

	_, err := s.HttpClient.R().
		SetResult(&customers).
		Get("/customer")

	if err != nil {
		return nil, err
	}

	return customers, nil
}

func (s *Service) GetByCpf(cpf string) (*Customer, error) {
	var customer *Customer

	_, err := s.HttpClient.R().
		SetPathParams(map[string]string{
			"cpf": cpf,
		}).
		SetResult(&customer).
		Get("/customer/{cpf}")

	if err != nil {
		return nil, err
	}

	return customer, nil

}
