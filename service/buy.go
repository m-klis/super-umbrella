package service

import (
	"gochicoba/models"
	"gochicoba/repository"
)

type BuyService interface {
	GetAllBuys() ([]*models.BuyResponse, error)
	CreateBuy(*models.Transaction) (*models.Transaction, error)
	CreateTransaction(*models.Transaction) (*models.Transaction, error)
}

type buyService struct {
	buyRepo repository.BuyRepository
}

func NewBuyService(buyRepo repository.BuyRepository) BuyService {
	return &buyService{
		buyRepo: buyRepo,
	}
}

func (bs *buyService) GetAllBuys() ([]*models.BuyResponse, error) {
	list, err := bs.buyRepo.GetAllBuys()
	if err != nil {
		return nil, err
	}
	return list, nil
}

func (bs *buyService) CreateBuy(t *models.Transaction) (*models.Transaction, error) {
	mt, err := bs.buyRepo.CreateBuy(t)
	if err != nil {
		return nil, err
	}
	return mt, nil
}

func (bs *buyService) CreateTransaction(t *models.Transaction) (*models.Transaction, error) {
	mt, err := bs.buyRepo.CreateTransaction(t)
	if err != nil {
		return nil, err
	}
	return mt, nil
}
