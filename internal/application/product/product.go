package product

import (
	"time"

	"github.com/google/uuid"
)

type Product struct {
	Id          uuid.UUID
	Name        string
	Description string
	Category    string
	Price       float64
	CreatedAt   time.Time
}

type Repository interface {
	Search(search string, category string) (*[]Product, error)
	Delete(id uuid.UUID) error
}
