package cart

import "context"

type Cart struct {
	ID       string
	ClientID string
	products []Product
}

type Product struct {
	ClientID  string
	ProductID string
	Quantity  int
	Comments  string
}

type ICartRepository interface {
	Create(ctx context.Context) (*Cart, error)
	Get(ctx context.Context, clientID string) (*Cart, error)
}

type ICartProductRepository interface {
	Create(ctx context.Context, cartID string, product *Product) error
}

type IService interface {
	LoadCart(ctx context.Context, clientID string) (*Cart, error)
	AddProduct(ctx context.Context, product *Product) error
}
