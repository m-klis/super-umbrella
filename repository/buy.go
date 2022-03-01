package repository

import (
	"gochicoba/models"

	"gorm.io/gorm"
)

type BuyRepository interface {
	GetAllBuys() ([]*models.BuyResponse, error)
	CreateBuy(*models.Transaction) (*models.Transaction, error)
	CreateTransaction(*models.Transaction) (*models.Transaction, error)
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

func (ur *buyRepository) CreateBuy(t *models.Transaction) (*models.Transaction, error) {
	query := ur.db

	err := query.Create(&t).Error
	if err != nil {
		return nil, err
	}

	return t, nil
}

func (ur *buyRepository) CreateTransaction(t *models.Transaction) (*models.Transaction, error) {
	query := ur.db

	err := query.Create(&t).Error
	if err != nil {
		return nil, err
	}

	return t, nil
}
