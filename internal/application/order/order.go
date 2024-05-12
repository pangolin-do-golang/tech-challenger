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

type Repository interface {
	Get(id uuid.UUID) (*Order, error)
}
