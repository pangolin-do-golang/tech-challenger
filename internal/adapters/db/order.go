package db

import (
	"fmt"

	"github.com/google/uuid"
	"github.com/pangolin-do-golang/tech-challenge/internal/application/order"
	"gorm.io/gorm"
)

type PostgresOrderRepository struct {
	db *gorm.DB
}

type OrderPostgres struct {
	BaseModel
	TotalAmount float64 `gorm:"total_amount"`
	Status      string  `gorm:"status"`
}

func (op OrderPostgres) TableName() string {
	return "order"
}

func NewPostgresOrderRepository(db *gorm.DB) order.Repository {
	return &PostgresOrderRepository{db: db}
}

func (repo *PostgresOrderRepository) Get(id uuid.UUID) (*order.Order, error) {
	var dbOrder OrderPostgres

	if err := repo.db.First(&dbOrder, "id = ?", id).Error; err != nil {

		fmt.Println("Order not found:", err)
		return nil, err
	}

	return &order.Order{
		Id:          dbOrder.ID,
		CreatedAt:   dbOrder.CreatedAt,
		TotalAmount: dbOrder.TotalAmount,
		Status:      dbOrder.Status,
	}, nil
}
