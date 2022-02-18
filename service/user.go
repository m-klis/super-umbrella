package service

import (
	"gochicoba/models"
	"gochicoba/repository"
)

type BuyService interface {
	GetAllBuys() ([]*models.BuyResponse, error)
	CreateBuy() (*models.TransactionResponse, error)
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

func (bs *buyService) CreateBuy() (*models.TransactionResponse, error) {
	list, err := bs.buyRepo.CreateBuy()
	if err != nil {
		return nil, err
	}
	return list, nil
}
