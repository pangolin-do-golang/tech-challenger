package order

import (
	"context"
	"time"

	"github.com/google/uuid"
)

const (
	StatusPending   = "PENDING"
	StatusCreated   = "CREATED"
	StatusPreparing = "PREPARING"
	StatusReady     = "READY"
	StatusFinished  = "FINISHED"
	StatusPaid      = "PAID"
	StatusDeclined  = "DECLINED"
)

type Order struct {
	ID          uuid.UUID `json:"id"`
	CreatedAt   time.Time `json:"created_at"`
	ClientID    uuid.UUID `json:"client_id"`
	TotalAmount float64   `json:"total_amount"`
	Status      string    `json:"status"`
}

type Product struct {
	ClientID  uuid.UUID `json:"client_id"`
	ProductID uuid.UUID `json:"product_id"`
	Quantity  int       `json:"quantity"`
	Comments  string    `json:"comments"`
	Total     float64   `json:"total"`
}

var Status = struct {
	Pending   string
	Created   string
	Preparing string
}{
	Created:   "CREATED",
	Pending:   "PENDING",
	Preparing: "PREPARING",
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
