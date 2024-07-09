package db

import (
	"context"

	"github.com/google/uuid"
	"github.com/pangolin-do-golang/tech-challenge/internal/core/cart"
	"gorm.io/gorm"
)

type PostgresCartProductsRepository struct {
	db *gorm.DB
}

type CartProductsPostgres struct {
	BaseModel
	CartID    uuid.UUID `gorm:"type:uuid" json:"cart_id"`
	ProductID uuid.UUID `gorm:"type:uuid" json:"product_id"`
	Quantity  int       `gorm:"quantity"`
	Comments  string    `gorm:"comments"`
}

func (op *CartProductsPostgres) TableName() string {
	return "cart_products"
}

func NewPostgresCartProductsRepository(db *gorm.DB) cart.ICartProductRepository {
	return &PostgresCartProductsRepository{db: db}
}

func (p *PostgresCartProductsRepository) Create(_ context.Context, cartID uuid.UUID, product *cart.Product) error {
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

func (p *PostgresCartProductsRepository) GetByCartID(_ context.Context, cartID uuid.UUID) ([]*cart.Product, error) {
	var cartProducts []CartProductsPostgres
	err := p.db.Where("cart_id = ?", cartID).Find(&cartProducts).Error
	if err != nil {
		return nil, err
	}

	var products []*cart.Product
	for _, cp := range cartProducts {
		products = append(products, &cart.Product{
			ProductID: cp.ProductID,
			Quantity:  cp.Quantity,
			Comments:  cp.Comments,
		})
	}

	return products, nil
}

func (p *PostgresCartProductsRepository) DeleteByProductID(_ context.Context, cartID, productID uuid.UUID) error {
	return p.db.Delete(&CartProductsPostgres{}, "cart_id = ? AND product_id = ?", cartID, productID).Error
}

func (p *PostgresCartProductsRepository) UpdateProductByProductID(_ context.Context, cartID, productID uuid.UUID, product *cart.Product) error {
	return p.db.Model(&CartProductsPostgres{}).
		Where("cart_id = ? AND product_id = ?", cartID, productID).
		Updates(map[string]interface{}{
			"quantity": product.Quantity,
			"comments": product.Comments,
		}).Error
}
