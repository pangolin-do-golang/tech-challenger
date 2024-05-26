package order

import (
	"context"
	"time"

	"github.com/google/uuid"
)

type Order struct {
	ID          uuid.UUID
	CreatedAt   time.Time
	ClientID    uuid.UUID
	TotalAmount float64
	Status      string
}

type Product struct {
	ClientID  uuid.UUID
	ProductID uuid.UUID
	Quantity  int
	Comments  string
	Total     float64
}

var Status = struct {
	Pending string
	Created string
}{
	Created: "CREATED",
	Pending: "PENDING",
}

type IOrderService interface {
	Get(id uuid.UUID) (*Order, error)
	Create(clientID uuid.UUID) (*Order, error)
	GetAll() ([]Order, error)
}

type IOrderProductRepository interface {
	Create(ctx context.Context, orderID uuid.UUID, product *Product) error
	GetByOrderID(ctx context.Context, orderID uuid.UUID) ([]*Product, error)
}

type IOrderRepository interface {
	Create(order *Order) (*Order, error)
	Update(order *Order) error
	Get(id uuid.UUID) (*Order, error)
	GetAll() ([]Order, error)
}
