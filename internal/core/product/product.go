package product

import (
	"time"

	"github.com/google/uuid"
)

type Product struct {
	Id          uuid.UUID `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Category    string    `json:"category"`
	Price       float64   `json:"price"`
	CreatedAt   time.Time `json:"created_at"`
}

type Repository interface {
	Search(search string, category string) (*[]Product, error)
	Delete(id uuid.UUID) error
	GetByID(id uuid.UUID) (*Product, error)
}
