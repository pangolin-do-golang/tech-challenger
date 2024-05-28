package db

import (
	"errors"

	"github.com/google/uuid"
	"github.com/pangolin-do-golang/tech-challenge/internal/application/cart"
	"github.com/pangolin-do-golang/tech-challenge/internal/domainerrors"
	"gorm.io/gorm"
)

type PostgresCartRepository struct {
	db *gorm.DB
}

type CartPostgres struct {
	BaseModel
	ClientID uuid.UUID              `gorm:"client_id"`
	Customer CustomerPostgres       `gorm:"foreignKey:ClientID"`
	Products []CartProductsPostgres `gorm:"foreignKey:CartID"`
}

func (op CartPostgres) TableName() string {
	return "cart"
}

func NewPostgresCartRepository(db *gorm.DB) cart.ICartRepository {
	return &PostgresCartRepository{db: db}
}

func (p *PostgresCartRepository) Create(clientID uuid.UUID) (*cart.Cart, error) {
	record := &CartPostgres{
		ClientID: clientID,
	}

	err := p.db.Create(record).Error
	if err != nil {
		return nil, err
	}

	return &cart.Cart{
		ID:       record.ID,
		ClientID: record.ClientID,
	}, nil
}

func (p *PostgresCartRepository) Get(clientID uuid.UUID) (*cart.Cart, error) {
	var row CartPostgres
	err := p.db.First(&row, "client_id = ?", clientID).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, domainerrors.ErrRecordNotFound
		}

		return nil, err
	}

	return &cart.Cart{
		ID:       row.ID,
		ClientID: row.ClientID,
	}, nil
}
