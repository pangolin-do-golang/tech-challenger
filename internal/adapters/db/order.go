package db

import (
	"github.com/google/uuid"
	"github.com/pangolin-do-golang/tech-challenge/internal/core/order"
	"github.com/pangolin-do-golang/tech-challenge/internal/errutil"
	"gorm.io/gorm"
)

type PostgresOrderRepository struct {
	db *gorm.DB
}

type OrderPostgres struct {
	BaseModel
	ClientID    uuid.UUID              `gorm:"client_id,type:uuid"`
	TotalAmount float64                `gorm:"total_amount"`
	Status      string                 `gorm:"status"`
	Customer    CustomerPostgres       `gorm:"foreignKey:ClientID"`
	Products    []OrderProductPostgres `gorm:"foreignKey:OrderID"`
}

func (op OrderPostgres) TableName() string {
	return "order"
}

func NewPostgresOrderRepository(db *gorm.DB) order.IOrderRepository {
	return &PostgresOrderRepository{db: db}
}

func (r *PostgresOrderRepository) Update(order *order.Order) error {
	result := r.db.Model(&OrderPostgres{}).
		Where("id", order.ID).
		Update("status", order.Status).
		Update("total_amount", order.TotalAmount)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (r *PostgresOrderRepository) Create(order *order.Order) (*order.Order, error) {
	dbOrder := OrderPostgres{
		BaseModel:   BaseModel{ID: uuid.New()},
		ClientID:    order.ClientID,
		TotalAmount: order.TotalAmount,
		Status:      order.Status,
	}

	result := r.db.Create(&dbOrder)
	if result.Error != nil {
		return nil, result.Error
	}

	order.ID = dbOrder.ID
	order.CreatedAt = dbOrder.CreatedAt
	return order, nil
}

func (r *PostgresOrderRepository) Get(id uuid.UUID) (*order.Order, error) {
	var record OrderPostgres

	if err := r.db.First(&record, "id = ?", id).Error; err != nil {
		return nil, errutil.ErrRecordNotFound
	}

	return &order.Order{
		ID:          record.ID,
		ClientID:    record.ClientID,
		CreatedAt:   record.CreatedAt,
		TotalAmount: record.TotalAmount,
		Status:      record.Status,
	}, nil
}

func (r *PostgresOrderRepository) GetAll() ([]order.Order, error) {
	var records []OrderPostgres

	err := r.db.Find(&records).Error
	if err != nil {
		return nil, err
	}

	parsedOrders := make([]order.Order, len(records))
	for i, record := range records {
		parsedOrders[i] = order.Order{
			ID:          record.ID,
			ClientID:    record.ClientID,
			CreatedAt:   record.CreatedAt,
			TotalAmount: record.TotalAmount,
			Status:      record.Status,
		}
	}

	return parsedOrders, nil
}
