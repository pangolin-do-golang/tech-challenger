package order

import (
	"time"

	"github.com/google/uuid"
)

type Order struct {
	Id          uuid.UUID
	CreatedAt   time.Time
	TotalAmount float64
	Status      string
}

type IService interface {
	Get(id uuid.UUID) (*Order, error)
	GetAll() ([]Order, error)
}

type IRepository interface {
	Get(id uuid.UUID) (*Order, error)
	GetAll() ([]Order, error)
}
