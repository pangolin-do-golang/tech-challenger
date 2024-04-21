package order

import "time"

type Order struct {
	Id          string
	CreatedAt   time.Time
	TotalAmount float64
	Status      string
}

type Repository interface {
	Get(id string) (*Order, error)
}
