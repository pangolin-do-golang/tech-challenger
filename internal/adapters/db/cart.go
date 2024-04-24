package db

import (
	"errors"
	"github.com/google/uuid"
	"github.com/pangolin-do-golang/tech-challenge/internal/application/cart"
	"github.com/pangolin-do-golang/tech-challenge/internal/domainerrors"
	"gorm.io/gorm"
	"time"
)

type PostgresCartRepository struct {
	db *gorm.DB
}

func (p *PostgresCartRepository) Create(clientID string) (*cart.Cart, error) {
	record := &CartPostgres{
		ID:       uuid.New().String(),
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

func (p *PostgresCartRepository) Get(clientID string) (*cart.Cart, error) {
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

type CartPostgres struct {
	ID        string    `gorm:"id"`
	ClientID  string    `gorm:"client_id"`
	CreatedAt time.Time `gorm:"created_at"`
}

func (op CartPostgres) TableName() string {
	return "cart"
}

func NewPostgresCartRepository(db *gorm.DB) cart.ICartRepository {
	return &PostgresCartRepository{db: db}
}
