package db

import (
	"fmt"
	"github.com/pangolin-do-golang/tech-challenge/internal/application/order"
	"gorm.io/gorm"
	"time"
)

type PostgresOrderRepository struct {
	db *gorm.DB
}

type OrderPostgres struct {
	Id          string    `gorm:"id"`
	CreatedAt   time.Time `gorm:"created_at"`
	TotalAmount float64   `gorm:"total_amount"`
	Status      string    `gorm:"status"`
}

func (op OrderPostgres) TableName() string {
	return "order"
}

func NewPostgresOrderRepository(db *gorm.DB) order.Repository {
	return &PostgresOrderRepository{db: db}
}

func (repo *PostgresOrderRepository) Get(id string) (*order.Order, error) {
	var dbOrder OrderPostgres

	if err := repo.db.First(&dbOrder, "id = ?", id).Error; err != nil {

		fmt.Println("Order not found:", err)
		return nil, err
	}

	return &order.Order{
		Id:          dbOrder.Id,
		CreatedAt:   dbOrder.CreatedAt,
		TotalAmount: dbOrder.TotalAmount,
		Status:      dbOrder.Status,
	}, nil
}
