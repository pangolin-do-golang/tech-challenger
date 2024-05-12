package cart

import (
	"context"

	"github.com/google/uuid"
)

type Cart struct {
	ID       uuid.UUID
	ClientID uuid.UUID
	products []Product
}

type Product struct {
	ClientID  uuid.UUID
	ProductID uuid.UUID
	Quantity  int
	Comments  string
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
	AddProduct(ctx context.Context, product *Product) error
	RemoveProduct(ctx context.Context, clientID uuid.UUID, productID uuid.UUID) error
	EditProduct(ctx context.Context, product *Product) error
}
