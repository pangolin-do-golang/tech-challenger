package db

import (
	"context"
	"github.com/pangolin-do-golang/tech-challenge/internal/application/cart"
	"gorm.io/gorm"
)

type PostgresCartProductsRepository struct {
	db *gorm.DB
}

func (p PostgresCartProductsRepository) Create(ctx context.Context, cartID string, product *cart.Product) error {
	cartProduct := CartProductsPostgres{
		CartID:    cartID,
		ProductID: product.ProductID,
		Quantity:  product.Quantity,
		Comments:  product.Comments,
	}

	result := p.db.Create(&cartProduct)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

type CartProductsPostgres struct {
	ID        string `gorm:"id"`
	CartID    string `gorm:"cart_id"`
	ProductID string `gorm:"product_id"`
	Quantity  int    `gorm:"quantity"`
	Comments  string `gorm:"comments"`
}

func (op CartProductsPostgres) TableName() string {
	return "cart_products"
}

func NewPostgresCartProductsRepository(db *gorm.DB) cart.ICartProductRepository {
	return &PostgresCartProductsRepository{db: db}
}
