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
	Create(clientID string) (*Cart, error)
	Get(clientID string) (*Cart, error)
}

type ICartProductRepository interface {
	Create(ctx context.Context, cartID string, product *Product) error
	GetByCartID(ctx context.Context, cartID string) ([]*Product, error)
	DeleteByProductID(ctx context.Context, cartID, productID string) error
	UpdateProductByProductID(ctx context.Context, cartID, productID string, product *Product) error
}

type IService interface {
	LoadCart(ctx context.Context, clientID string) (*Cart, error)
	AddProduct(ctx context.Context, product *Product) error
	RemoveProduct(ctx context.Context, clientID string, productID string) error
	EditProduct(ctx context.Context, product *Product) error
}
