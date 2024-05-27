package db

import (
	"errors"
	"fmt"

	"github.com/google/uuid"
	"github.com/pangolin-do-golang/tech-challenge/internal/application/order"
	"github.com/pangolin-do-golang/tech-challenge/internal/domainerrors"
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

func NewPostgresOrderRepository(db *gorm.DB) order.IOrderRepository {
	return &PostgresOrderRepository{db: db}
}

func (r *PostgresOrderRepository) Update(order *order.Order) error {
	dbOrder := OrderPostgres{
		BaseModel:   BaseModel{ID: order.ID},
		TotalAmount: order.TotalAmount,
		Status:      order.Status,
	}

	result := r.db.Save(&dbOrder)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (r *PostgresOrderRepository) Create(order *order.Order) (*order.Order, error) {
	dbOrder := OrderPostgres{
		BaseModel:   BaseModel{ID: order.ID},
		TotalAmount: order.TotalAmount,
		Status:      order.Status,
	}

	result := r.db.Create(&dbOrder)
	if result.Error != nil {
		return nil, result.Error
	}

	order.ID = dbOrder.ID
	return order, nil
}

func (r *PostgresOrderRepository) Get(id uuid.UUID) (*order.Order, error) {
	var record OrderPostgres

	if err := r.db.First(&record, "id = ?", id).Error; err != nil {

		fmt.Println("Order not found:", err)
		return nil, err
	}

	return &order.Order{
		ID:          record.ID,
		CreatedAt:   record.CreatedAt,
		TotalAmount: record.TotalAmount,
		Status:      record.Status,
	}, nil
}

func (r *PostgresOrderRepository) GetAll() ([]order.Order, error) {
	var records []OrderPostgres

	err := r.db.Find(&records).Error

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, domainerrors.ErrRecordNotFound
	}

	if err != nil {
		return nil, err
	}

	parsedOrders := make([]order.Order, len(records))
	for i, record := range records {
		parsedOrders[i] = order.Order{
			ID:          record.ID,
			CreatedAt:   record.CreatedAt,
			TotalAmount: record.TotalAmount,
			Status:      record.Status,
		}
	}

	return parsedOrders, nil
}
