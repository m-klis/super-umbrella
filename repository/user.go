package repository

import (
	"gochicoba/models"

	"gorm.io/gorm"
)

type BuyRepository interface {
	GetAllBuys() ([]*models.BuyResponse, error)
	CreateBuy() (*models.TransactionResponse, error)
}

type buyRepository struct {
	db *gorm.DB
}

func NewBuyRepository(db *gorm.DB) BuyRepository {
	return &buyRepository{
		db: db,
	}
}

func (ur *buyRepository) GetAllBuys() ([]*models.BuyResponse, error) {
	return nil, nil
}

func (ur *buyRepository) CreateBuy() (*models.TransactionResponse, error) {
	return nil, nil
}
