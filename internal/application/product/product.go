package product

import "time"

type Product struct {
	Id          string
	Name        string
	Description string
	Category    string
	Price       float64
	CreatedAt   time.Time
}

type Repository interface {
	Search(search string, category string) (*[]Product, error)
	Delete(id string) error
}
