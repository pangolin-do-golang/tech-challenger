package db

import (
	"errors"
	"fmt"
	"strings"

	"github.com/google/uuid"
	"github.com/pangolin-do-golang/tech-challenge/internal/core/order"
	"github.com/pangolin-do-golang/tech-challenge/internal/domainerrors"
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
	Products    []OrderProductPostgres `gorm:"foreignKey:OrderID"`
	Customer    CustomerPostgres       `gorm:"foreignKey:ClientID"`
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
		ClientID:    order.ClientID,
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

		fmt.Println("Order not found:", err)
		return nil, err
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

	err := r.db.Raw(buildGetAllQuery()).Scan(&records).Error

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
			ClientID:    record.ClientID,
			CreatedAt:   record.CreatedAt,
			TotalAmount: record.TotalAmount,
			Status:      record.Status,
		}
	}

	return parsedOrders, nil
}

func buildGetAllQuery() string {
	ignoredStatus := []string{
		order.StatusFinished,
		order.StatusDeclined,
	}

	statusSortedPriority := []string{
		order.StatusReady,
		order.StatusPreparing,
		order.StatusPaid,
		order.StatusPending,
		order.StatusCreated,
	}

	return fmt.Sprintf(`
		SELECT * FROM "order"
		WHERE status NOT IN (%s)
		ORDER BY
			CASE %s
			END, created_at
		`,
		buildIgnoredStatusCondition(ignoredStatus),
		buildSortedStatusCase(statusSortedPriority),
	)
}

func buildIgnoredStatusCondition(ignoredStatus []string) string {
	var ignoredParts []string
	for _, status := range ignoredStatus {
		ignoredParts = append(ignoredParts, fmt.Sprintf("'%s'", status))
	}

	return strings.Join(ignoredParts, ", ")
}

func buildSortedStatusCase(statusSortedPriority []string) string {
	var caseParts []string
	for i, status := range statusSortedPriority {
		caseParts = append(caseParts, fmt.Sprintf("WHEN status = '%s' THEN %d", status, i+1))
	}
	caseParts = append(caseParts, fmt.Sprintf("ELSE %d", len(statusSortedPriority)))

	return strings.Join(caseParts, " ")
}
