package db

import (
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/pangolin-do-golang/tech-challenge/internal/application/product"
	"gorm.io/gorm"
)

type PostgresProductRepository struct {
	db *gorm.DB
}

type ProductPostgres struct {
	BaseModel
	Name        string  `gorm:"name"`
	Description string  `gorm:"description"`
	Category    string  `gorm:"category"`
	Price       float64 `gorm:"price"`
}

func (pp ProductPostgres) TableName() string {
	return "product"
}

func NewPostgresProductRepository(db *gorm.DB) product.Repository {
	return &PostgresProductRepository{db: db}
}

func (repo *PostgresProductRepository) GetByID(id uuid.UUID) (*product.Product, error) {
	var dbRecord ProductPostgres

	if err := repo.db.Where("id = ? and deleted_at is null", id).First(&dbRecord).Error; err != nil {
		return nil, err
	}

	if dbRecord.ID == uuid.Nil {
		return nil, nil
	}

	return &product.Product{
		Id:          dbRecord.ID,
		Name:        dbRecord.Name,
		Description: dbRecord.Description,
		Category:    dbRecord.Category,
		Price:       dbRecord.Price,
	}, nil
}

func (repo *PostgresProductRepository) Search(search string, category string) (*[]product.Product, error) {
	var dbRecords []ProductPostgres

	db := repo.db

	if len(search) > 0 {
		parsedSearch := fmt.Sprintf("%%%s%%", search)
		db = db.Where("name ILIKE ?", parsedSearch)
	}

	if len(category) > 0 {
		db = db.Where("category = ?", category)
	}

	if err := db.Where("deleted_at is null").Find(&dbRecords).Error; err != nil {
		return nil, err
	}

	if len(dbRecords) == 0 {
		emptyResult := make([]product.Product, 0)
		return &emptyResult, nil
	}

	var results []product.Product

	for _, record := range dbRecords {
		results = append(results, product.Product{
			Id:          record.ID,
			Name:        record.Name,
			Description: record.Description,
			Category:    record.Category,
			Price:       record.Price,
		})
	}

	return &results, nil
}

func (repo *PostgresProductRepository) Delete(id uuid.UUID) error {
	var dbRecord *ProductPostgres

	if err := repo.db.Where("id = ? and deleted_at is null", id).First(&dbRecord).Error; err != nil {
		return err
	}

	if dbRecord == nil {
		return nil
	}

	if err := repo.db.Model(&dbRecord).Where("id = ? and deleted_at is null", id).Update("deleted_at", time.Now()).Error; err != nil {
		fmt.Println(err)
		return err
	}

	return nil
}
