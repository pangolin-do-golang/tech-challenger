package cart

import (
	"context"

	"github.com/google/uuid"
)

type Cart struct {
	ID       uuid.UUID  `json:"id"`
	ClientID uuid.UUID  `json:"client_id"`
	Products []*Product `json:"products"`
}

type Product struct {
	ClientID  uuid.UUID `json:"client_id"`
	ProductID uuid.UUID `json:"product_id"`
	Quantity  int       `json:"quantity"`
	Comments  string    `json:"comments"`
}

type ICartRepository interface {
	Create(clientID uuid.UUID) (*Cart, error)
	Get(clientID uuid.UUID) (*Cart, error)
}

type ICartProductRepository interface {
	Create(ctx context.Context, cartID uuid.UUID, product *Product) error
	GetByCartID(ctx context.Context, cartID uuid.UUID) ([]*Product, error)
	DeleteByProductID(ctx context.Context, cartID, productID uuid.UUID) error
	UpdateProductByProductID(ctx context.Context, cartID, productID uuid.UUID, product *Product) error
}

type IService interface {
	LoadCart(ctx context.Context, clientID uuid.UUID) (*Cart, error)
	GetFullCart(clientID uuid.UUID) (*Cart, error)
	AddProduct(ctx context.Context, product *Product) error
	RemoveProduct(ctx context.Context, clientID uuid.UUID, productID uuid.UUID) error
	EditProduct(ctx context.Context, product *Product) error
	Cleanup(clientID uuid.UUID) error
}
