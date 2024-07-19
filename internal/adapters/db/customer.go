package db

import (
	"github.com/google/uuid"
	"github.com/pangolin-do-golang/tech-challenge/internal/core/customer"
	"github.com/pangolin-do-golang/tech-challenge/internal/domainerrors"
	"gorm.io/gorm"
)

type PostgresCustomerRepository struct {
	db *gorm.DB
}

type CustomerPostgres struct {
	BaseModel
	Name  string `gorm:"name"`
	Cpf   string `gorm:"uniqueIndex" json:"cpf"`
	Email string `gorm:"email"`
	Age   int    `gorm:"age"`
}

func (cp CustomerPostgres) TableName() string {
	return "customer"
}

func NewPostgresCustomerRepository(db *gorm.DB) customer.IRepository {
	return &PostgresCustomerRepository{db: db}
}

func (r *PostgresCustomerRepository) Create(cust customer.Customer) (*customer.Customer, error) {
	record := &CustomerPostgres{
		Name:  cust.Name,
		Cpf:   cust.Cpf,
		Email: cust.Email,
		Age:   cust.Age,
	}

	if err := r.db.Create(&record).Error; err != nil {
		return nil, err
	}

	return &customer.Customer{
		Id:    record.ID,
		Name:  record.Name,
		Cpf:   record.Cpf,
		Email: record.Email,
		Age:   record.Age,
	}, nil
}

func (r *PostgresCustomerRepository) Update(customerId uuid.UUID, cust customer.Customer) (*customer.Customer, error) {
	var record CustomerPostgres
	err := r.db.First(&record, customerId).Error

	if domainerrors.Is(err, gorm.ErrRecordNotFound) {
		return nil, domainerrors.ErrRecordNotFound
	}

	record = CustomerPostgres{
		Name:  cust.Name,
		Cpf:   cust.Cpf,
		Email: cust.Email,
		Age:   cust.Age,
	}

	record.ID = customerId

	if err := r.db.Save(&record).Error; err != nil {
		return nil, err
	}

	return &customer.Customer{
		Id:    record.ID,
		Name:  record.Name,
		Cpf:   record.Cpf,
		Email: record.Email,
		Age:   record.Age,
	}, nil
}

func (r *PostgresCustomerRepository) Delete(customerId uuid.UUID) error {
	var record CustomerPostgres
	err := r.db.Delete(&record, customerId).Error

	if domainerrors.Is(err, gorm.ErrRecordNotFound) {
		return domainerrors.ErrRecordNotFound
	}

	if err != nil {
		return err
	}

	return nil
}

func (r *PostgresCustomerRepository) GetAll() ([]customer.Customer, error) {
	var records []CustomerPostgres

	err := r.db.Find(&records).Error

	if domainerrors.Is(err, gorm.ErrRecordNotFound) {
		return nil, domainerrors.ErrRecordNotFound
	}

	if err != nil {
		return nil, err
	}

	parsedCustomers := make([]customer.Customer, len(records))
	for i, record := range records {
		parsedCustomers[i] = customer.Customer{
			Id:    record.ID,
			Name:  record.Name,
			Cpf:   record.Cpf,
			Email: record.Email,
			Age:   record.Age,
		}
	}

	return parsedCustomers, nil
}

func (r *PostgresCustomerRepository) GetByCpf(customerCpf string) (*customer.Customer, error) {
	var record CustomerPostgres

	err := r.db.Where("cpf = ?", customerCpf).First(&record).Error

	if domainerrors.Is(err, gorm.ErrRecordNotFound) {
		return nil, domainerrors.ErrRecordNotFound
	}

	if err != nil {
		return nil, err
	}

	return &customer.Customer{
		Id:    record.ID,
		Name:  record.Name,
		Cpf:   record.Cpf,
		Email: record.Email,
		Age:   record.Age,
	}, nil
}
